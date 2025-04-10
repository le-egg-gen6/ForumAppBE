#!/bin/bash
set -e

# Create replication user
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $REPLICATION_USER WITH REPLICATION PASSWORD '$REPLICATION_PASSWORD';
EOSQL

# Modify pg_hba.conf to allow replication connections
cat >> "$PGDATA/pg_hba.conf" <<EOF
# Replication connections
host replication $REPLICATION_USER all md5
host all all all md5
EOF

# Configure postgresql.conf for replication
cat >> "$PGDATA/postgresql.conf" <<EOF
# Replication settings
listen_addresses = '*'
wal_level = replica
max_wal_senders = 10
wal_keep_size = '1GB'
hot_standby = on
archive_mode = on
archive_command = 'cp %p $PGDATA/archive/%f'
EOF

# Create archive directory
mkdir -p "$PGDATA/archive"

# Restart PostgreSQL to apply changes
pg_ctl -D "$PGDATA" -m fast -w restart

# Create a database to test replication
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE TABLE IF NOT EXISTS test_replication (id SERIAL PRIMARY KEY, data TEXT);
  INSERT INTO test_replication (data) VALUES ('Initial master data');
EOSQL

echo "Master node initialized successfully"
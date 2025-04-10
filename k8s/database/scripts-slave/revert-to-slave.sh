#!/bin/bash
set -e

echo "Reverting to slave mode..."

# Stop PostgreSQL to reconfigure
pg_ctl -D "$PGDATA" -m fast -w stop || true

# Clear data directory
rm -rf "$PGDATA"/*

# Use pg_basebackup to clone the master
until pg_basebackup -h $MASTER_HOST -D "$PGDATA" -U $REPLICATION_USER -P -v --wal-method=stream; do
    echo "Waiting for master to become available..."
    sleep 5
done

# Reconfigure as standby
cat > "$PGDATA/standby.signal" <<EOF
# This file indicates that this PostgreSQL instance is a standby
EOF

cat >> "$PGDATA/postgresql.conf" <<EOF
# Standby configuration
primary_conninfo = 'host=$MASTER_HOST port=5432 user=$REPLICATION_USER password=$REPLICATION_PASSWORD application_name=slave'
hot_standby = on
EOF

# Start PostgreSQL again
pg_ctl -D "$PGDATA" -w start

echo "Successfully reverted to slave mode"

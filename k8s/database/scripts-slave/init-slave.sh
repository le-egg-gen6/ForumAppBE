#!/bin/bash
set -e

# Check if PostgreSQL data directory is empty
if [ -s "$PGDATA/PG_VERSION" ]; then
    echo "PostgreSQL data directory is not empty, skipping slave initialization"
    exit 0
fi

echo "Initializing slave node..."

# Stop PostgreSQL server if running
pg_ctl -D "$PGDATA" -m fast -w stop || true

# Clear data directory
rm -rf "$PGDATA"/*

# Use pg_basebackup to clone the master
until pg_basebackup -h $MASTER_HOST -D "$PGDATA" -U $REPLICATION_USER -P -v --wal-method=stream; do
    echo "Waiting for master to become available..."
    sleep 5
done

# Configure recovery.conf settings
cat > "$PGDATA/standby.signal" <<EOF
# This file indicates that this PostgreSQL instance is a standby
EOF

cat >> "$PGDATA/postgresql.conf" <<EOF
# Standby configuration
primary_conninfo = 'host=$MASTER_HOST port=5432 user=$REPLICATION_USER password=$REPLICATION_PASSWORD application_name=slave'
hot_standby = on
EOF

echo "Slave node initialized successfully"
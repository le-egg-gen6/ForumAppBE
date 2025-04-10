#!/bin/bash
set -e

echo "Promoting slave to master..."
pg_ctl promote -D "$PGDATA"

# Update postgresql.conf to reflect master status
cat >> "$PGDATA/postgresql.conf" <<EOF
# Now operating as master
EOF

echo "Slave promoted to master successfully"
#!/bin/bash
set -e

pg_isready -U $POSTGRES_USER || exit 1

# Check if this is a master node
if psql -U $POSTGRES_USER -c "SELECT pg_is_in_recovery()" -t | grep -q f; then
    echo "Node is running as master"
    exit 0
else
    echo "Node should be master but is in recovery mode"
    exit 1
fi

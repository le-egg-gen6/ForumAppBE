#!/bin/bash

# Function to check if master is available
check_master() {
    pg_isready -h $MASTER_HOST -p 5432 -U $POSTGRES_USER >/dev/null 2>&1
}

# Function to check if we are currently the master
is_master() {
    ! psql -U $POSTGRES_USER -c "SELECT pg_is_in_recovery()" -t | grep -q t
}

# Main monitoring loop
while true; do
    # If master is down and we're still a slave, promote to master
    if ! check_master && ! is_master; then
        echo "Master is down, promoting this slave to master..."
        /usr/local/bin/promote-to-master.sh
    fi

    # If master is back up and we're running as master, revert to slave
    if check_master && is_master; then
        echo "Original master is back, reverting to slave..."
        /usr/local/bin/revert-to-slave.sh
    fi

    # Wait before next check
    sleep 10
done
#!/bin/bash
set -e

echo "Stopping PostgreSQL database service..."
cd ./database
docker-compose down
echo "PostgreSQL database service stopped."
cd ..

echo "Stopping Redis service..."
cd ./rediz
docker-compose down
echo "Redis service stopped."
cd ..

echo "All services have been stopped."
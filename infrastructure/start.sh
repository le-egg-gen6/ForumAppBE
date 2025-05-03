#!/bin/bash
set -e

echo "Starting Redis service..."
cd ./rediz
docker-compose up -d
echo "Redis service started successfully."
cd ..

echo "Starting PostgreSQL database service..."
cd ./database
docker-compose up -d
echo "PostgreSQL database service started successfully."
cd ..

echo "All services are running! Your setup includes:"
echo "- Redis at localhost:9999 (password: forum)"
echo "- PostgreSQL at localhost:6789 (username: forum, password: forum, database: forum)"

# Optional: Display running containers
echo ""
echo "Running containers:"
docker ps
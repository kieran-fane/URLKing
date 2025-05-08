#!/bin/bash

# Start Docker in background
docker compose up -d

# Trap SIGINT and SIGTERM to shut down containers
trap 'echo "Stopping..."; docker compose down; exit' INT TE

go run server.go

docker compose down


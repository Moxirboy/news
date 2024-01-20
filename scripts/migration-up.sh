#!/bin/bash

# Source environment variables from the .env file
if [ -f .env ]; then
  source .env
fi

# Run docker-compose command
 docker compose run --rm migrate -path=migrations/ -database="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable" up

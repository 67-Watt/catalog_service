#!/bin/bash
# Script to run database migrations

DB_URL=${DB_URL:-"postgres://user:password@localhost:5432/catalog_db?sslmode=disable"}

echo "Running migrations on $DB_URL..."

migrate -path ./migrations -database "$DB_URL" up

if [ $? -eq 0 ]; then
  echo "Migrations ran successfully."
else
  echo "Migration failed."
  exit 1
fi

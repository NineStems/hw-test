#!/bin/bash

# Collect static files
echo "getting goose"
go install github.com/pressly/goose/v3/cmd/goose@latest

# Apply database migrations
echo "set up migrations"
goose -dir="migrations" postgres "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable" down

echo "starting service"
./calendar --config=configs/calendar_config
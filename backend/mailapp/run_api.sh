#!/bin/bash

# Run the go application locally with a localhost db config
# Set the environment variables
ZINC_DB_URL=http://localhost:4080
ZINC_DB_USER=admin
ZINC_DB_PASSWORD=Complexpass#123
export ZINC_DB_URL ZINC_DB_USER ZINC_DB_PASSWORD

# Run the Go application
exec go run main.go
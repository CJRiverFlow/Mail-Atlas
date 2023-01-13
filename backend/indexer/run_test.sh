#!/bin/bash

# Run the indexer process with data available in local
# Set the environment variables
ZINC_DB_URL=http://localhost:4080
ZINC_DB_USER=admin
ZINC_DB_PASSWORD=Complexpass#123
export ZINC_DB_URL ZINC_DB_USER ZINC_DB_PASSWORD

#execute the data loading
exec go test
#!/bin/sh


# exit immediately if a command return a !=0 status
set -e  
echo "run migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "start the app"
# take all param (CMD) passed into and run it
exec "$@"

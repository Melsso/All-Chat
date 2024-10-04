#!/bin/sh

SESSION_KEY=$(openssl rand -base64 32)


export SESSION_KEY=$SESSION_KEY

while true; do 
    nc -z db 3306
    if [ $? -eq 0 ]; then
        echo "Database is ready!"
        break 
    else
        echo "Waiting for db..."
        sleep 1
    fi
done

exec ./main

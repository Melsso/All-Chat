#!/bin/sh

# Generate a random session key
SESSION_KEY=$(openssl rand -base64 32)

# Write the session key to the .env file
cat <<EOF > .env
SESSION_KEY=$SESSION_KEY
DB_HOST=db
DB_PORT=3306
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydb
EOF

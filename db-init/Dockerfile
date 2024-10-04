FROM mariadb:latest

# Set environment variables for the root password and database creation
ENV MYSQL_ROOT_PASSWORD=Melsopvvp_007@
ENV MYSQL_DATABASE=chatdb
ENV MYSQL_USER=serv
ENV MYSQL_PASSWORD=pswd

# Copy your initialization script to the right directory
COPY ./init.sql /docker-entrypoint-initdb.d/

# Expose the default port for MariaDB
EXPOSE 3306

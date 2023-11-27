FROM postgres:latest
ENV POSTGRES_USER=psqlman
ENV POSTGRES_PASSWORD=9cAN5M6J4x9M5Ly7qC
ENV POSTGRES_DB=TodoDB

# Copy the initialization script to create the default user
COPY init.sql /docker-entrypoint-initdb.d/
#COPY init.sql /docker-entrypoint-initdb.d/
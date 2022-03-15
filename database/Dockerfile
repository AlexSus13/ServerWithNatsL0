FROM postgres:12.9-alpine
ENV POSTGRES_USER alexey
ENV POSTGRES_DB ordersdb 
ENV POSTGRES_PASSWORD 12345
ADD createTable.sql /docker-entrypoint-initdb.d/

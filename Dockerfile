
FROM postgres:16.1-alpine3.19

# Set environment variables
ENV POSTGRES_USER=vladik
ENV POSTGRES_PASSWORD=123456
ENV POSTGRES_DB=students

# Expose the PostgreSQL port
EXPOSE 5432
COPY init.sql /docker-entrypoint-initdb.d/
# Run the container in detached mode
CMD ["postgres"]

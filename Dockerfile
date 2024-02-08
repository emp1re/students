##docker run --rm --name postgres -e POSTGRES_USER=vladik -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=students -p 5432:5432 -d postgres:latest
FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=vladik
ENV POSTGRES_PASSWORD=123456
ENV POSTGRES_DB=students

# Expose the PostgreSQL port
EXPOSE 5432
COPY init.sql /docker-entrypoint-initdb.d/
# Run the container in detached mode
CMD ["postgres"]
postgres:
	docker run --rm --name postgres -e POSTGRES_USER=vladik -e POSTGRES_PASSWORD=123456  -p 5432:5432 -d postgres:16.1-alpine3.19
createdb:
	docker exec -it postgres createdb --username=vladik --owner=vladik students
dropdb:
	docker exec -it postgres dropdb --username=vladik students
migrateup:
	migrate -database "postgresql://vladik:123456@localhost:5432/students?sslmode=disable" -verbose -path pkg/db/migrate up
migratedown:
	migrate -database "postgresql://vladik:123456@localhost:5432/students?sslmode=disable" -verbose -path pkg/db/migrate down
migrateforce1:
	migrate -database "postgresql://vladik:123456@localhost:5432/students?sslmode=disable" -path pkg/db/migrate force 1
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateup migratedown migrateforce1 sqlc
DB_URL=postgresql://root:secret@localhost:5434/pfdb?sslmode=disable

postgres:
	# docker run --name postgres12 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
	docker run --name ps12 -p 5434:5432 -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root pfdb

ds:
	docker start postgres12

dropdb:
	docker exec -it postgres12 pfdb

migrateup:
	go run api/migrations/migrate.go up

migratedown:
	go run api/migrations/migrate.go down

server:
	go run main.go

test:
	DATABASE_CONNECTION="user=postgres password=secret host=localhost port=5434 sslmode=disable" go test ./... -v --cover

.PHONY: postgres createdb dropdb server test ds

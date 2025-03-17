# docker contaner for postgres image
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# remove the postgres container
rmpostgres:
	docker stop postgres12 && docker rm postgres12

# create a database in the postgres container
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

# drop the database in the postgres container
dropdb:
	docker exec -it postgres12 dropdb simple_bank

# migrate the database in the postgres container
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

# migrate down the database in the postgres container
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# generate the go file(sqlc) from the queries
sqlc:
	sqlc generate

# run test for the project
test:
	go test -v -cover ./...

# format the project files
fmt:
	go fmt ./...

.PHONY: postgres rmpostgres createdb dropdb migrateup migratedown
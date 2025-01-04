docker:
	docker run --name power -p 0001:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

sqlc:
	sqlc generate

mgup:
	migrate -database "postgres://root:secret@localhost:0001/postgres?sslmode=disable" -path ./migrations up

mgdown:
	migrate -database "postgres://root:secret@localhost:0001/postgres?sslmode=disable" -path ./migrations down -all


phony : sqlc, docker
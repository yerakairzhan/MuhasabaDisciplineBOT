docker:
	docker run --name muhasaba -p 0365:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

sqlc:
	sqlc generate

mgup:
	migrate -database "postgres://root:secret@localhost:0365/postgres?sslmode=disable" -path ./migrations up

mgdown:
	migrate -database "postgres://root:secret@localhost:0365/postgres?sslmode=disable" -path ./migrations down

phony : sqlc, docker
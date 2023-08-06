postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -d postgres:15.3

database:
	docker exec -it postgres15 psql -U postgres

runDB:
	docker start postgres15

migrateUp:
	migrate -path gormDB/migration -database "postgres://postgres:1234@localhost:5432/gogin?sslmode=disable" -verbose up

migrateDown:
	migrate -path gormDB/migration -database "postgres://postgres:1234@localhost:5432/gogin?sslmode=disable" -verbose down

.PHONY: postgres database migrateUp migrateDown

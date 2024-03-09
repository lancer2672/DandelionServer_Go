postgres:
	docker run --name dandelion  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=JBdragonfire1135 -d postgres
createdb:
	docker exec -it dandelion createdb --username=root --owner=root dandelion_go
dropdb:
	docker exec -it dandelion dropdb dandelion_go
migrateup:
	 migrate -path db/migration -database "postgresql://root:JBdragonfire1135@localhost:5432/dandelion_go?sslmode=disable" -verbose up
migrateup1:
	 migrate -path db/migration -database "postgresql://root:JBdragonfire1135@localhost:5432/dandelion_go?sslmode=disable" -verbose up 1
migratedown:
	 migrate -path db/migration -database "postgresql://root:JBdragonfire1135@localhost:5432/dandelion_go?sslmode=disable" -verbose down
migratedown1:
	 migrate -path db/migration -database "postgresql://root:JBdragonfire1135@localhost:5432/dandelion_go?sslmode=disable" -verbose down 1
sqlc: 
	sqlc generate
server: 
	go run main.go
.PHONY: server postgres createdb dropdb migrateup migratedown sqlc migrateup1 migratedown1
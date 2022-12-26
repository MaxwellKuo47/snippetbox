postgres:
	docker run --name postgres12-shop-dev -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12-shop-dev createdb --username=root --owner=root shop_dev 

dropdb:
	docker exec -it postgres12-shop-dev dropdb shop_dev

migrateup:
	migrate -path internal/models/migration -database "postgresql://root:secret@localhost:5432/shop_dev?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/models/migration -database "postgresql://root:secret@localhost:5432/shop_dev?sslmode=disable" -verbose down

server:
	go run ./cmd/web

.PHONY: postgres createdb dropdb migrateup migratedown server
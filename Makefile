pull:
	docker pull mysql:5.7.40

mysql:
	docker run --name mysql-letsGo -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:5.7.40

rmmysql:
	docker stop mysql-letsGo
	docker rm mysql-letsGo

createdb:
	docker exec -it mysql-letsGo sh -c "mysql -u root -p'secret' -e 'CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;'"

dropdb:
	docker exec -it mysql-letsGo sh -c "mysql -u root -p'secret' -e 'DROP DATABASE snippetbox;'"

showdb:
	docker exec -it mysql-letsGo sh -c "mysql -u root -p'secret' -e 'SHOW DATABASES;'"

migrateup:
	migrate -path internal/models/migration -database "mysql://root:secret@tcp(localhost:3306)/snippetbox?x-tls-insecure-skip-verify=false" -verbose up

migratedown:
	migrate -path internal/models/migration -database "mysql://root:secret@tcp(localhost:3306)/snippetbox?x-tls-insecure-skip-verify=false" -verbose down

server:
	go run ./cmd/web

.PHONY: pull mysql rmmysql createdb dropdb migrateup migratedown server
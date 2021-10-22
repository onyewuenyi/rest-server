postgres:
	docker run --name taskstore -e POSTGRES_USER=charlesonyewuenyi -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres postgres -N 1000


createdb:
	winpty docker exec -it taskstore createdb --username=charlesonyewuenyi taskstore 

dropdb:
	winpty docker exec -it taskstore dropdb --username=charlesonyewuenyi taskstore

migrateup:
	migrate -path db/migration -database "postgres://charlesonyewuenyi:password@localhost:5432/taskstore?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://charlesonyewuenyi:password@localhost:5432/taskstore?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
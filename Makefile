postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root lcrm

run_local:
	docker run --name lcrm-back -p 8080:8080  -d lcrm-backend:1.0

sqlc:
	sqlc generate
.PHONY: postgres createdb deletedb migrateup migratedown sqlc



c_m:
	migrate create -ext sql -dir db/migrations -seq $(name)

p_up:
	docker-compose up -d

p_down:
	docker-compose down

db_up:
	docker exec -it bank-postgres createdb --username=root --owner=root bank_db

db_down:
	docker exec -it bank-postgres dropdb --username=root bank_db

m_up:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/bank_db?sslmode=disable" up

m_down:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5432/bank_db?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

start:
	CompileDaemon -command="./bank_db
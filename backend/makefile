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
build:
	docker-compose build

run:
	docker-compose up

migrate:
	migrate -path ./migrations -database 'mysql://root:root@tcp(localhost:3306)/golang_example' up
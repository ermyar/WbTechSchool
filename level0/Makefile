build:
	docker compose build 'app'

rebuild:
	docker stop level0-app && docker rm level0-app
	make run

postgres:
	docker compose up postgres

# 	psql -h localhost -p 5433 -U admin -d db   to connect to postgres

run:
	docker compose -f 'docker-compose.yaml' up --build 'app'

go-producer: 
	go run ./cmd/producer/main.go

kafka:
	docker compose -f 'docker-compose.yaml' up -d kafka-ui

clean:
	docker compose down
	docker image prune -f

.PHONY: clean
build:
	docker compose build 'app'

run:
	docker compose -f 'docker-compose.yaml' up -d --build 'app'

go-producer: 
	go run ./cmd/producer/main.go

kafka:
	docker compose -f 'docker-compose.yaml' up -d kafka-ui

clean:
	docker compose down
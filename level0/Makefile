build:
	docker compose build 'app'

run:
	docker compose -f 'docker-compose.yaml' up -d --build 'app'

run-producer:
	go run ./cmd/producer/main.go
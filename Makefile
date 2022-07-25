.PHONY: seed dc-up docker build

seed:
	docker exec -i boletia-db /bin/bash -c "PGPASSWORD=password psql --username postgres boletia" < ./migrations/01_currency.up.sql

dc-up:
	docker compose -f docker-compose.yml up -d

dc-down:
	docker compose -f docker-compose.yml down

docker:
	docker build -t golang/air .

build:
	go build
.PHONY: seed dc-up

seed:
	docker exec -i boletia-db /bin/bash -c "PGPASSWORD=password psql --username postgres boletia" < ./migrations/01_currency.up.sql

dc-up:
	docker-compose -f docker-compose.yml up -d

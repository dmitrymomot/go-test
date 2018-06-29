.PHONY: run down restart pma

run:
	docker-compose up -d db \
	&& sleep 10 \
	&& docker-compose up -d api

down:
	docker-compose down
	docker rmi -f testapi

restart: down run

pma:
	docker-compose up -d pma

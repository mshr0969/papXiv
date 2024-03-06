.PHONY: gen
gen:
	oapi-codegen  -package handler --config configs/openapi/config.yaml ../spec/docs/openapi.yaml > internal/handler/openapi.gen.go

.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down

.PHONY: docker-clean
docker-clean:
	docker rm -f `docker ps -a -q`
	docker volume rm `docker volume ls -q`

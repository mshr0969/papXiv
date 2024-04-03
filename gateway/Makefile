.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: gen
gen:
	oapi-codegen  -package handler --config configs/openapi/config.yaml ../spec/docs/openapi.yaml > internal/handler/openapi.gen.go

.PHONY: up
up:
	docker compose up -d
	sleep 20
	make migrate-up

.PHONY: down
down:
	docker compose down

.PHONY: docker-clean
docker-clean:
	docker rm -f `docker ps -a -q`
	docker volume rm `docker volume ls -q`

.PHONY: test
test:
	go test ./...

MIGRATE_OPT_DEV ?= -env development -config ./configs/migrations/config/dbconfig.yml

.PHONY: migrate-new
migrate-new:
	set -a; source .env.local; set +a; sql-migrate new $(MIGRATE_OPT_DEV) ${NAME}

.PHONY: migrate-up
migrate-up:
	set -a; source .env.local; set +a; sql-migrate up $(MIGRATE_OPT_DEV)

.PHONY: migrate-status
migrate-status:
	set -a; source .env.local; set +a; sql-migrate status $(MIGRATE_OPT_DEV)

.PHONY: migrate-down
migrate-down:
	set -a; source .env.local; set +a; sql-migrate down $(MIGRATE_OPT_DEV)

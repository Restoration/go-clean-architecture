ifneq (,$(wildcard ./.env))
	include ./.env
endif

TEST_POSTGRES_DATABASE = "test"

DSN=host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DATABASE) port=$(POSTGRES_PORT) sslmode=disable
CMD=docker-compose exec -e GOOSE_DBSTRING="$(DSN)" app

DSN1=host=$(POSTGRES_HOST_1) user=$(POSTGRES_USER_1) password=$(POSTGRES_PASSWORD_1) dbname=$(POSTGRES_DATABASE_1) port=$(POSTGRES_PORT_1) sslmode=disable
CMD1=docker-compose exec -e GOOSE_DBSTRING="$(DSN1)" app

DSN2=host=$(POSTGRES_HOST_2) user=$(POSTGRES_USER_2) password=$(POSTGRES_PASSWORD_2) dbname=$(POSTGRES_DATABASE_2) port=$(POSTGRES_PORT_2) sslmode=disable
CMD2=docker-compose exec -e GOOSE_DBSTRING="$(DSN2)" app

TEST_DSN=host=$(POSTGRES_HOST) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(TEST_POSTGRES_DATABASE) port=$(POSTGRES_PORT) sslmode=disable
CMD_TEST=docker-compose exec -e GOOSE_DBSTRING="$(TEST_DSN)" app

.PHONY: init
init:
	cp ./dotenv .env

create_migrate:
	$(CMD) goose -dir infrastructure/db/migrations create ${FILE} sql

create_seed:
	$(CMD) goose -dir infrastructure/db/seeds create ${FILE} sql

migrate_up:
	$(CMD) goose -allow-missing -dir infrastructure/db/migrations postgres up

migrate_up_all:
	$(CMD) goose -allow-missing -dir infrastructure/db/migrations postgres up
	$(CMD1) goose -allow-missing -dir infrastructure/db/migrations postgres up
	$(CMD2) goose -allow-missing -dir infrastructure/db/migrations postgres up

migrate_down:
	$(CMD) goose -dir infrastructure/db/migrations postgres down

migrate_reset:
	$(CMD) goose -dir infrastructure/db/seeds -table goose_seed_version postgres reset
	$(CMD) goose -dir infrastructure/db/migrations postgres reset

seed:
	$(CMD) goose -dir infrastructure/db/seeds -table goose_seed_version postgres up

seed_reset:
	$(CMD) goose -dir infrastructure/db/seeds -table goose_seed_version postgres reset


unit_test:
	$(CMD_TEST) go test -v ./test/units/application/interactor
	$(CMD_TEST) go test -v ./test/units/presentation/controller

e2e_test:
	$(CMD_TEST) goose -allow-missing -dir infrastructure/db/migrations postgres up
	$(CMD_TEST) goose -dir infrastructure/db/seeds -table goose_seed_version postgres up
	$(CMD_TEST) go test -v ./test/e2e
	$(CMD_TEST) goose -dir infrastructure/db/seeds -table goose_seed_version postgres reset
	$(CMD_TEST) goose -dir infrastructure/db/migrations postgres reset

.PHONY: sqlc, dev

dev:
	@docker compose up --watch

sqlc:
	@sqlc generate -f ./api/db/sqlc.yaml --no-remote

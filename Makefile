.PHONY: sqlc

ifndef-env:
	@[ -f .env ] || cp .env.example .env

# --- Primary Targets ---

dev: ifndef-env
	docker compose up --watch

sqlc:
	@sqlc generate -f ./api/db/sqlc.yaml --no-remote

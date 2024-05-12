sql-up:
	sql-migrate up DB_HOST=localhost DB_PORT=5432 DB_NAME=viska_db DB_USER=postgres DB_PASSWORD=password

sql-down:
	sql-migrate down -env="dev" DB_HOST=localhost DB_PORT=5432 DB_NAME=viska_db DB_USER=postgres DB_PASSWORD=password

db-up:
	@docker compose -f ./docker-compose.yml up -d db

db-down:
	@docker compose -f ./docker-compose.yml down db

compose-up:
	@docker-compose up -d

compose-down:
	@docker-compose down

docs:
	swag fmt && swag init
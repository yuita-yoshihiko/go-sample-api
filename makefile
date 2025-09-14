db-up:
	docker compose up -d db

db-down:
	docker compose down db

migration/new:
	go tool sql-migrate new --env="local" ${FILE_NAME}

migrate/up:
	make db-up
	sleep 5
	go tool sql-migrate up --env="local"

migrate/down:
	make db-up
	sleep 5
	go tool sql-migrate down --env="local"

db-seed:
	docker compose exec -T db sh -c "psql -v ON_ERROR_STOP=1 postgres://postgres:password@db:5432/go-sample-api_local?sslmode=disable" < ./initdb/initdb.sql

# テスト用DBコンテナ立ち上げ
test-db-up:
	docker compose -f docker-compose.testdb.yml up --renew-anon-volumes -d --wait
	${RUN} sh -c "go tool sql-migrate up --env='test'"

# テスト用DBコンテナ落とす
test-db-down:
	docker compose -f docker-compose.testdb.yml down --volumes

# ローカルDB接続
psql:
	docker compose exec db psql -U postgres -d go-sample-api_local

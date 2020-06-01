api = server

# APIコンテナに接続
login:
	docker-compose exec $(api) sh

# magration
migrate:
	docker-compose exec $(api) sh -c 'goose up'

# rollback
rollback:
	docker-compose exec $(api) sh -c 'goose down'

# Mysqlに作業ユーザーで接続
mysql:
	docker-compose exec db bash -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'

# テスト実行
test:
	go test -v ./infrastructure/persistence
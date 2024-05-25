# 環境変数の設定
DB_USER = imaikosuke
DB_NAME = plog

# SQLスクリプトのパス
RESET_DB_SQL = migrations/reset_db.sql
CREATE_TABLES_SQL = migrations/create_tables.sql

# Goのメインファイルのパス
MAIN_GO = cmd/main/main.go

.PHONY: run reset create start

# データベースのリセット
reset:
	@echo "Resetting database..."
	@psql -U $(DB_USER) -d $(DB_NAME) -f $(RESET_DB_SQL)

# テーブルの作成
create:
	@echo "Creating tables..."
	@psql -U $(DB_USER) -d $(DB_NAME) -f $(CREATE_TABLES_SQL)

# サーバーの起動
start:
	@echo "Starting server..."
	@go run $(MAIN_GO)

# 全てのコマンドを実行
run: reset create start

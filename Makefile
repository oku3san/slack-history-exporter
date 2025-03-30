.PHONY: build test clean lint

# 変数定義
BINARY_NAME=slack-history-exporter
BUILD_DIR=bin
MAIN_PATH=cmd/slack-history-exporter/main.go

# ビルド
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# テスト実行
test:
	@echo "Running tests..."
	@go test -v ./...

# リント実行
lint:
	@echo "Running linter..."
	@go fmt ./...
	@go vet ./...

# クリーン
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean

# 依存関係のインストール
deps:
	@echo "Installing dependencies..."
	@go mod tidy

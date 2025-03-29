# Slack履歴エクスポートツール用Makefile

# 変数
BINARY_NAME=slack-history-exporter
GO=go
GOFMT=gofmt
GOFILES=$(shell find . -name "*.go" -type f)
GOPATH=$(shell $(GO) env GOPATH)
VERSION=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

# デフォルトターゲット
.PHONY: all
all: fmt test build

# ビルド
.PHONY: build
build:
	@echo "ビルドを開始します..."
	$(GO) build $(LDFLAGS) -o $(BINARY_NAME) ./cmd/slack-history-exporter

# テスト
.PHONY: test
test:
	@echo "テストを実行します..."
	$(GO) test -v ./...

# コードフォーマット
.PHONY: fmt
fmt:
	@echo "コードをフォーマットします..."
	$(GOFMT) -w $(GOFILES)

# 依存関係の更新
.PHONY: deps
deps:
	@echo "依存関係を更新します..."
	$(GO) mod tidy

# クリーン
.PHONY: clean
clean:
	@echo "生成されたファイルを削除します..."
	rm -f $(BINARY_NAME)
	rm -f *.json *.csv

# インストール
.PHONY: install
install: build
	@echo "バイナリをインストールします..."
	cp $(BINARY_NAME) $(GOPATH)/bin/

# ヘルプ
.PHONY: help
help:
	@echo "利用可能なコマンド:"
	@echo "  make          : コードをフォーマットし、テストを実行し、ビルドします"
	@echo "  make build    : バイナリをビルドします"
	@echo "  make test     : テストを実行します"
	@echo "  make fmt      : コードをフォーマットします"
	@echo "  make deps     : 依存関係を更新します"
	@echo "  make clean    : 生成されたファイルを削除します"
	@echo "  make install  : バイナリをインストールします"
	@echo "  make help     : このヘルプを表示します"
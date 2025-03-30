# Slack履歴エクスポートツール

## 概要

Slack履歴エクスポートツールは、Slack上の会話履歴をエクスポートするためのコマンドラインツールです。指定されたSlackチャンネルからメッセージ履歴を取得し、ローカルにJSON形式またはCSV形式でエクスポートできます。

## 主要機能

- Slackのチャンネル履歴をエクスポート
- コマンドライン引数による操作
- 環境変数によるユーザートークン認証
- JSON/CSV形式でのエクスポート
- スレッド返信の取得
- 日時指定による履歴取得

## 開発ロードマップ

プロジェクトは以下のタスクに分割して開発を進めます：

1. [プロジェクト初期設定](issues/01_project_setup.md)
2. [Slack API認証機能の実装](issues/02_authentication.md)
3. [Slackチャンネル履歴取得機能の実装](issues/03_history_retrieval.md)
4. [Slack履歴エクスポート機能の実装](issues/04_export_functionality.md)
5. [エラーハンドリング機能の実装](issues/05_error_handling.md)
6. [コマンドライン引数処理機能の実装](issues/06_command_line_args.md)
7. [統合テストとエンドツーエンドテストの実装](issues/07_integration_testing.md)
8. [ドキュメント作成](issues/08_documentation.md)

## 開発の進め方

本プロジェクトはTest Driven Development（TDD）アプローチで開発を進めます：

1. 各機能の要件を理解する
2. テストを先に書く
3. テストが失敗することを確認する
4. 機能を実装する
5. テストが成功することを確認する
6. コードをリファクタリングする

## 環境設定

### 必要条件

- Go 1.24以上
- Slackユーザートークン

### インストール方法

```bash
# リポジトリのクローン
git clone https://github.com/yourusername/slack-history-exporter.git
cd slack-history-exporter

# 依存関係のインストール
go mod download

# ビルド
make build
```

## 使用方法

```bash
# 環境変数の設定
export SLACK_TOKEN="xoxp-xxxxxxxxxx-xxxxxxxxxx-xxxxxxxxxxxx"

# 基本的な使用法
./slack-history-exporter -c C12345678

# 出力先を指定
./slack-history-exporter -c C12345678 -o my-export.json

# CSVフォーマットでエクスポート
./slack-history-exporter -c C12345678 -f csv
```

## 貢献方法

1. このリポジトリをフォークする
2. 新しいブランチを作成する (`git checkout -b feature/amazing-feature`)
3. 変更をコミットする (`git commit -m 'Add some amazing feature'`)
4. ブランチにプッシュする (`git push origin feature/amazing-feature`)
5. プルリクエストを作成する

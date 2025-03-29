# Slack履歴エクスポートツール

Slack上の会話履歴をエクスポートするためのコマンドラインツールです。指定されたSlackチャンネルからメッセージ履歴を取得し、JSONまたはCSV形式でローカルにエクスポートできます。

## 機能

- Slackのチャンネル履歴をエクスポート
- JSONまたはCSV形式でのエクスポート
- スレッド返信の取得
- 日数指定による履歴取得
- 詳細ログの出力

## インストール

### 前提条件

- Go 1.24以上

### インストール方法

```bash
go install github.com/oku3san/slack-history-exporter@latest
```

または、リポジトリをクローンしてビルドする方法：

```bash
git clone https://github.com/oku3san/slack-history-exporter.git
cd slack-history-exporter
go build -o slack-history-exporter ./cmd/slack-history-exporter
```

## 使用方法

### 環境変数の設定

Slackのユーザートークンを環境変数に設定します：

```bash
export SLACK_TOKEN="xoxp-xxxxxxxxxx-xxxxxxxxxx-xxxxxxxxxxxx"
```

### 基本的な使用法

```bash
slack-history-exporter -c C12345678
```

### オプション

- `--channel`, `-c`: エクスポート対象のチャンネルID (必須)
- `--output`, `-o`: 出力ファイル名 (デフォルト: `{チャンネル名}_{日付}.json`)
- `--format`, `-f`: 出力フォーマット (デフォルト: `json`、オプション: `csv`)
- `--days`, `-d`: 取得する日数 (デフォルト: 全履歴)
- `--include-threads`, `-i`: スレッド返信も取得するかどうか (デフォルト: `true`)
- `--verbose`, `-v`: 詳細ログの出力
- `--help`, `-h`: ヘルプ表示

### 使用例

```bash
# 基本的な使用法
slack-history-exporter -c C12345678

# 出力先を指定
slack-history-exporter -c C12345678 -o my-export.json

# 最近7日分のみエクスポート
slack-history-exporter -c C12345678 -d 7

# CSVフォーマットでエクスポート
slack-history-exporter -c C12345678 -f csv

# 詳細ログを出力
slack-history-exporter -c C12345678 -v
```

## 開発

### テスト実行

```bash
go test ./...
```

### 依存関係

- [github.com/slack-go/slack](https://github.com/slack-go/slack) - Slack APIクライアント
- [github.com/spf13/cobra](https://github.com/spf13/cobra) - コマンドライン引数解析
- [github.com/pkg/errors](https://github.com/pkg/errors) - エラーハンドリング

## ライセンス

MIT

## 注意事項

- Slack APIの制限により大量のリクエストには制限がかかる場合があります
- ユーザートークンは個人のアクセス権限に依存するため、アクセスできないチャンネルの履歴は取得できません
- ファイル添付のダウンロードは現在サポートされていません

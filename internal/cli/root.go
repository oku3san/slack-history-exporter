package cli

import (
	"fmt"
	"github.com/oku3san/slack-history-exporter/internal/api"
	"github.com/oku3san/slack-history-exporter/internal/exporter"
	"os"

	"github.com/spf13/cobra"
)

var (
	// コマンドラインフラグ
	channelID      string
	outputFile     string
	outputFormat   string
	days           int
	includeThreads bool
	verbose        bool

	// ルートコマンド
	rootCmd = &cobra.Command{
		Use:   "slack-history-exporter",
		Short: "Slackのチャンネル履歴をエクスポートするツール",
		Long: `slack-history-exporterは、指定されたSlackチャンネルからメッセージ履歴を取得し、
ローカルにエクスポートするためのコマンドラインツールです。`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 環境変数からSlackトークンを取得
			token := os.Getenv("SLACK_TOKEN")
			if token == "" {
				return fmt.Errorf("SLACK_TOKEN環境変数が設定されていません")
			}

			// チャンネルIDが指定されているか確認
			if channelID == "" {
				return fmt.Errorf("チャンネルIDを指定してください (--channel または -c フラグを使用)")
			}

			// エクスポート処理の実行
			return runExport(token, channelID, outputFile, outputFormat, days, includeThreads, verbose)
		},
	}
)

// Execute はルートコマンドを実行します
func Execute() error {
	return rootCmd.Execute()
}

// runExport はエクスポート処理を実行します
func runExport(token, channelID, outputFile, outputFormat string, days int, includeThreads, verbose bool) error {
	// Slack APIクライアントの作成
	client := api.NewClient(token, verbose)

	// チャンネル履歴の取得
	if verbose {
		fmt.Printf("チャンネル %s の履歴を取得中...\n", channelID)
	}

	data, err := client.GetChannelHistory(channelID, days, includeThreads)
	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("%d件のメッセージを取得しました\n", len(data.Messages))
	}

	// エクスポーターの作成
	exp, err := exporter.NewExporter(outputFormat)
	if err != nil {
		return fmt.Errorf("エクスポーターの作成に失敗しました: %w", err)
	}

	// エクスポート
	if verbose {
		fmt.Printf("%s形式でエクスポート中...\n", outputFormat)
	}

	return exp.Export(data, outputFile)
}

func init() {
	// フラグの定義
	rootCmd.PersistentFlags().StringVarP(&channelID, "channel", "c", "", "エクスポート対象のチャンネルID (必須)")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "出力ファイル名 (デフォルト: {チャンネル名}_{日付}.json)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "f", "json", "出力フォーマット (json または csv)")
	rootCmd.PersistentFlags().IntVarP(&days, "days", "d", 0, "取得する日数 (デフォルト: 全履歴)")
	rootCmd.PersistentFlags().BoolVarP(&includeThreads, "include-threads", "i", true, "スレッド返信も取得するかどうか")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "詳細ログの出力")

	// チャンネルIDを必須に設定
	rootCmd.MarkPersistentFlagRequired("channel")
}

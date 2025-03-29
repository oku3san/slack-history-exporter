package exporter

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/oku3san/slack-history-exporter/internal/models"
)

// CSVExporter はCSVフォーマットでのエクスポートを行います
type CSVExporter struct{}

// NewCSVExporter は新しいCSVエクスポーターを作成します
func NewCSVExporter() *CSVExporter {
	return &CSVExporter{}
}

// Export はデータをCSV形式でエクスポートします
func (e *CSVExporter) Export(data *models.ExportData, outputFile string) error {
	// 出力ファイル名が指定されていない場合はデフォルト名を使用
	if outputFile == "" {
		outputFile = fmt.Sprintf("%s_%s.csv", data.ChannelName, time.Now().Format("2006-01-02"))
	} else if filepath.Ext(outputFile) == "" {
		// 拡張子がない場合は.csvを追加
		outputFile = outputFile + ".csv"
	}

	// ファイルを作成
	file, err := os.Create(outputFile)
	if err != nil {
		return errors.Wrap(err, "ファイルの作成に失敗しました")
	}
	defer file.Close()

	// CSVライターを作成
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// ヘッダーを書き込み
	header := []string{"メッセージID", "ユーザー", "ユーザー名", "テキスト", "タイムスタンプ", "スレッド"}
	if err := writer.Write(header); err != nil {
		return errors.Wrap(err, "ヘッダーの書き込みに失敗しました")
	}

	// メッセージを書き込み
	for _, msg := range data.Messages {
		// スレッド情報
		threadInfo := ""
		if len(msg.ThreadReplies) > 0 {
			threadInfo = fmt.Sprintf("%d件の返信", len(msg.ThreadReplies))
		}

		// メッセージ行を作成
		row := []string{
			msg.MessageID,
			msg.User,
			msg.UserName,
			msg.Text,
			msg.Timestamp.Format(time.RFC3339),
			threadInfo,
		}

		// 行を書き込み
		if err := writer.Write(row); err != nil {
			return errors.Wrap(err, "メッセージの書き込みに失敗しました")
		}

		// スレッド返信も書き込み
		for _, reply := range msg.ThreadReplies {
			replyRow := []string{
				reply.MessageID,
				reply.User,
				reply.UserName,
				fmt.Sprintf("    %s", reply.Text), // インデントを付けて返信であることを示す
				reply.Timestamp.Format(time.RFC3339),
				"",
			}

			if err := writer.Write(replyRow); err != nil {
				return errors.Wrap(err, "返信メッセージの書き込みに失敗しました")
			}
		}
	}

	fmt.Printf("エクスポート完了: %s\n", outputFile)
	return nil
}

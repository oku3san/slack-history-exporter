package exporter

import (
	"os"
	"testing"
	"time"

	"github.com/oku3san/slack-history-exporter/internal/models"
)

func TestNewExporter(t *testing.T) {
	// JSONエクスポーターのテスト
	jsonExporter, err := NewExporter("json")
	if err != nil {
		t.Fatalf("JSONエクスポーターの作成に失敗しました: %v", err)
	}
	if jsonExporter == nil {
		t.Fatal("JSONエクスポーターがnilです")
	}

	// CSVエクスポーターのテスト
	csvExporter, err := NewExporter("csv")
	if err != nil {
		t.Fatalf("CSVエクスポーターの作成に失敗しました: %v", err)
	}
	if csvExporter == nil {
		t.Fatal("CSVエクスポーターがnilです")
	}

	// 不正なフォーマットのテスト
	invalidExporter, err := NewExporter("invalid")
	if invalidExporter != nil {
		t.Error("不正なフォーマットでエクスポーターが作成されました")
	}
}

func TestJSONExporter(t *testing.T) {
	// テスト用のデータを作成
	now := time.Now()
	data := &models.ExportData{
		ChannelID:   "C12345678",
		ChannelName: "general",
		ExportDate:  now,
		Messages: []models.Message{
			{
				MessageID: "1234567890.123456",
				User:      "U12345678",
				UserName:  "testuser",
				Text:      "テストメッセージ",
				Timestamp: now,
			},
		},
	}

	// JSONエクスポーターを作成
	exporter := NewJSONExporter()

	// 一時ファイルにエクスポート
	tempFile := "test_export.json"
	defer os.Remove(tempFile) // テスト終了後にファイルを削除

	err := exporter.Export(data, tempFile)
	if err != nil {
		t.Fatalf("エクスポートに失敗しました: %v", err)
	}

	// ファイルが作成されたか確認
	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		t.Fatal("エクスポートファイルが作成されていません")
	}
}

func TestCSVExporter(t *testing.T) {
	// テスト用のデータを作成
	now := time.Now()
	data := &models.ExportData{
		ChannelID:   "C12345678",
		ChannelName: "general",
		ExportDate:  now,
		Messages: []models.Message{
			{
				MessageID: "1234567890.123456",
				User:      "U12345678",
				UserName:  "testuser",
				Text:      "テストメッセージ",
				Timestamp: now,
				ThreadReplies: []models.Message{
					{
						MessageID: "1234567890.123457",
						User:      "U87654321",
						UserName:  "replyuser",
						Text:      "返信メッセージ",
						Timestamp: now.Add(time.Hour),
					},
				},
			},
		},
	}

	// CSVエクスポーターを作成
	exporter := NewCSVExporter()

	// 一時ファイルにエクスポート
	tempFile := "test_export.csv"
	defer os.Remove(tempFile) // テスト終了後にファイルを削除

	err := exporter.Export(data, tempFile)
	if err != nil {
		t.Fatalf("エクスポートに失敗しました: %v", err)
	}

	// ファイルが作成されたか確認
	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		t.Fatal("エクスポートファイルが作成されていません")
	}
}

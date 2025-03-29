package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMessageJSON(t *testing.T) {
	// テスト用のメッセージを作成
	now := time.Now()
	msg := Message{
		MessageID: "1234567890.123456",
		User:      "U12345678",
		UserName:  "testuser",
		Text:      "テストメッセージ",
		Timestamp: now,
		Reactions: []Reaction{
			{
				Name:  "thumbsup",
				Count: 3,
				Users: []string{"U11111111", "U22222222", "U33333333"},
			},
		},
	}

	// JSONにエンコード
	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("JSONエンコードに失敗しました: %v", err)
	}

	// JSONからデコード
	var decodedMsg Message
	if err := json.Unmarshal(data, &decodedMsg); err != nil {
		t.Fatalf("JSONデコードに失敗しました: %v", err)
	}

	// 元のメッセージと一致するか確認
	if msg.MessageID != decodedMsg.MessageID {
		t.Errorf("MessageID: 期待値=%s, 実際=%s", msg.MessageID, decodedMsg.MessageID)
	}
	if msg.User != decodedMsg.User {
		t.Errorf("User: 期待値=%s, 実際=%s", msg.User, decodedMsg.User)
	}
	if msg.UserName != decodedMsg.UserName {
		t.Errorf("UserName: 期待値=%s, 実際=%s", msg.UserName, decodedMsg.UserName)
	}
	if msg.Text != decodedMsg.Text {
		t.Errorf("Text: 期待値=%s, 実際=%s", msg.Text, decodedMsg.Text)
	}
	if !msg.Timestamp.Equal(decodedMsg.Timestamp) {
		t.Errorf("Timestamp: 期待値=%v, 実際=%v", msg.Timestamp, decodedMsg.Timestamp)
	}
	if len(msg.Reactions) != len(decodedMsg.Reactions) {
		t.Errorf("Reactions長さ: 期待値=%d, 実際=%d", len(msg.Reactions), len(decodedMsg.Reactions))
	}
}

func TestExportDataJSON(t *testing.T) {
	// テスト用のエクスポートデータを作成
	now := time.Now()
	data := ExportData{
		ChannelID:   "C12345678",
		ChannelName: "general",
		ExportDate:  now,
		Messages:    []Message{},
	}

	// JSONにエンコード
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("JSONエンコードに失敗しました: %v", err)
	}

	// JSONからデコード
	var decodedData ExportData
	if err := json.Unmarshal(jsonData, &decodedData); err != nil {
		t.Fatalf("JSONデコードに失敗しました: %v", err)
	}

	// 元のデータと一致するか確認
	if data.ChannelID != decodedData.ChannelID {
		t.Errorf("ChannelID: 期待値=%s, 実際=%s", data.ChannelID, decodedData.ChannelID)
	}
	if data.ChannelName != decodedData.ChannelName {
		t.Errorf("ChannelName: 期待値=%s, 実際=%s", data.ChannelName, decodedData.ChannelName)
	}
	if !data.ExportDate.Equal(decodedData.ExportDate) {
		t.Errorf("ExportDate: 期待値=%v, 実際=%v", data.ExportDate, decodedData.ExportDate)
	}
}

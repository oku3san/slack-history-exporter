package api

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	// クライアントの作成テスト
	client := NewClient("dummy-token", false)
	if client == nil {
		t.Fatal("クライアントがnilです")
	}

	// verboseモードのクライアント作成テスト
	verboseClient := NewClient("dummy-token", true)
	if verboseClient == nil {
		t.Fatal("verboseクライアントがnilです")
	}
}

func TestParseTimestamp(t *testing.T) {
	// 正常なタイムスタンプのテスト
	ts, err := parseTimestamp("1609459200.123456")
	if err != nil {
		t.Fatalf("タイムスタンプの解析に失敗しました: %v", err)
	}

	// 期待値の確認（2021-01-01 00:00:00 UTC）
	expected := int64(1609459200)
	if ts.Unix() != expected {
		t.Errorf("タイムスタンプが一致しません: 期待値=%d, 実際=%d", expected, ts.Unix())
	}

	// 不正なタイムスタンプのテスト
	_, err = parseTimestamp("invalid")
	if err == nil {
		t.Error("不正なタイムスタンプでエラーが発生しませんでした")
	}
}

// 注意: 以下のテストは実際のSlack APIを呼び出すため、
// 実行時にはSLACK_TOKENが設定されている必要があります。
// また、実際のチャンネルIDが必要です。
// これらのテストは通常はスキップされます。

func TestGetChannelInfo_Skip(t *testing.T) {
	t.Skip("このテストは実際のSlack APIを呼び出すためスキップされます")

	client := NewClient("your-token-here", true)
	channel, err := client.GetChannelInfo("your-channel-id-here")
	if err != nil {
		t.Fatalf("チャンネル情報の取得に失敗しました: %v", err)
	}
	if channel == nil {
		t.Fatal("チャンネル情報がnilです")
	}
}

func TestGetChannelHistory_Skip(t *testing.T) {
	t.Skip("このテストは実際のSlack APIを呼び出すためスキップされます")

	client := NewClient("your-token-here", true)
	data, err := client.GetChannelHistory("your-channel-id-here", 1, true)
	if err != nil {
		t.Fatalf("チャンネル履歴の取得に失敗しました: %v", err)
	}
	if data == nil {
		t.Fatal("エクスポートデータがnilです")
	}
}

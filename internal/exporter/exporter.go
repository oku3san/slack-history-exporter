package exporter

import (
	"fmt"

	"github.com/oku3san/slack-history-exporter/pkg/api"
)

// Exporter はSlackのエクスポートを処理するための構造体です
type Exporter struct {
	client *api.Client
}

// NewExporter は新しいExporterを作成します
func NewExporter(token string) *Exporter {
	return &Exporter{
		client: api.NewClient(token),
	}
}

// ExportChannel は指定されたチャンネルの履歴をエクスポートします
func (e *Exporter) ExportChannel(channelID string, outputPath string) error {
	fmt.Printf("Exporting history for channel %s to %s\n", channelID, outputPath)

	// ここでチャンネル履歴の取得と保存の実装を行う
	// 実装はこれから

	return nil
}

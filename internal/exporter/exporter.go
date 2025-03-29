package exporter

import (
	"github.com/oku3san/slack-history-exporter/internal/models"
)

// Exporter はエクスポート機能を提供するインターフェースです
type Exporter interface {
	// Export はデータをエクスポートします
	Export(data *models.ExportData, outputFile string) error
}

// NewExporter は指定されたフォーマットに対応するエクスポーターを作成します
func NewExporter(format string) (Exporter, error) {
	switch format {
	case "json":
		return NewJSONExporter(), nil
	case "csv":
		return NewCSVExporter(), nil
	default:
		return nil, nil
	}
}

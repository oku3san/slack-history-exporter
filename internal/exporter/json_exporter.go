package exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/oku3san/slack-history-exporter/internal/models"
	"github.com/pkg/errors"
)

// JSONExporter はJSONフォーマットでのエクスポートを行います
type JSONExporter struct{}

// NewJSONExporter は新しいJSONエクスポーターを作成します
func NewJSONExporter() *JSONExporter {
	return &JSONExporter{}
}

// Export はデータをJSON形式でエクスポートします
func (e *JSONExporter) Export(data *models.ExportData, outputFile string) error {
	// 出力ファイル名が指定されていない場合はデフォルト名を使用
	if outputFile == "" {
		outputFile = fmt.Sprintf("%s_%s.json", data.ChannelName, time.Now().Format("2006-01-02"))
	} else if filepath.Ext(outputFile) == "" {
		// 拡張子がない場合は.jsonを追加
		outputFile = outputFile + ".json"
	}

	// JSONエンコード
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Wrap(err, "JSONエンコードに失敗しました")
	}

	// ファイルに書き込み
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		return errors.Wrap(err, "ファイルの書き込みに失敗しました")
	}

	fmt.Printf("エクスポート完了: %s\n", outputFile)
	return nil
}

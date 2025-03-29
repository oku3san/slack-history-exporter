package main

import (
	"fmt"
	"os"

	"github.com/oku3san/slack-history-exporter/internal/cli"
)

func main() {
	// コマンドの実行
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %s\n", err)
		os.Exit(1)
	}
}
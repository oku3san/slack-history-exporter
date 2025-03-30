package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slack-history-exporter",
	Short: "A tool to export Slack history",
	Long:  `A command line tool to export and archive Slack conversation history.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Slack History Exporter - Use subcommands to perform operations")
	},
}

func init() {
	// ここにサブコマンドを追加
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}

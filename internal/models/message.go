package models

import (
	"time"
)

// Message はSlackのメッセージを表す構造体です
type Message struct {
	MessageID     string     `json:"message_id"`
	User          string     `json:"user"`
	UserName      string     `json:"user_name"`
	Text          string     `json:"text"`
	Timestamp     time.Time  `json:"timestamp"`
	Reactions     []Reaction `json:"reactions,omitempty"`
	ThreadReplies []Message  `json:"thread_replies,omitempty"`
}

// Reaction はメッセージに対するリアクションを表す構造体です
type Reaction struct {
	Name  string   `json:"name"`
	Count int      `json:"count"`
	Users []string `json:"users,omitempty"`
}

// ExportData はエクスポートデータ全体を表す構造体です
type ExportData struct {
	ChannelID   string    `json:"channel_id"`
	ChannelName string    `json:"channel_name"`
	ExportDate  time.Time `json:"export_date"`
	Messages    []Message `json:"messages"`
}

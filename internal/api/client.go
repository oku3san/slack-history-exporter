package api

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/slack-go/slack"

	"github.com/oku3san/slack-history-exporter/internal/models"
)

// Client はSlack APIクライアントのラッパーです
type Client struct {
	api     *slack.Client
	verbose bool
}

// NewClient は新しいSlack APIクライアントを作成します
func NewClient(token string, verbose bool) *Client {
	return &Client{
		api:     slack.New(token),
		verbose: verbose,
	}
}

// GetChannelInfo はチャンネル情報を取得します
func (c *Client) GetChannelInfo(channelID string) (*slack.Channel, error) {
	channel, err := c.api.GetConversationInfo(&slack.GetConversationInfoInput{
		ChannelID: channelID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "チャンネル情報の取得に失敗しました")
	}
	return channel, nil
}

// GetChannelHistory はチャンネルの履歴を取得します
func (c *Client) GetChannelHistory(channelID string, days int, includeThreads bool) (*models.ExportData, error) {
	// チャンネル情報の取得
	channel, err := c.GetChannelInfo(channelID)
	if err != nil {
		return nil, err
	}

	// エクスポートデータの初期化
	exportData := &models.ExportData{
		ChannelID:   channelID,
		ChannelName: channel.Name,
		ExportDate:  time.Now(),
		Messages:    []models.Message{},
	}

	// 履歴取得のパラメータ設定
	params := slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Limit:     100, // 一度に取得するメッセージ数
	}

	// 日数が指定されている場合は、その日数分の履歴のみ取得
	if days > 0 {
		oldest := time.Now().AddDate(0, 0, -days).Unix()
		params.Oldest = fmt.Sprintf("%f", float64(oldest))
	}

	// 履歴の取得（ページネーションは簡略化）
	history, err := c.api.GetConversationHistory(&params)
	if err != nil {
		return nil, errors.Wrap(err, "チャンネル履歴の取得に失敗しました")
	}

	// メッセージの処理
	for _, msg := range history.Messages {
		message, err := c.convertToMessage(msg, channelID, includeThreads)
		if err != nil {
			if c.verbose {
				fmt.Printf("警告: メッセージの変換に失敗しました: %s\n", err)
			}
			continue
		}
		exportData.Messages = append(exportData.Messages, *message)
	}

	return exportData, nil
}

// convertToMessage はSlackのメッセージをモデルに変換します
func (c *Client) convertToMessage(msg slack.Message, channelID string, includeThreads bool) (*models.Message, error) {
	// タイムスタンプの解析
	ts, err := parseTimestamp(msg.Timestamp)
	if err != nil {
		return nil, errors.Wrap(err, "タイムスタンプの解析に失敗しました")
	}

	// メッセージの作成
	message := &models.Message{
		MessageID: msg.Timestamp,
		User:      msg.User,
		UserName:  c.getUserName(msg.User),
		Text:      msg.Text,
		Timestamp: ts,
		Reactions: []models.Reaction{},
	}

	// リアクションの処理
	for _, reaction := range msg.Reactions {
		message.Reactions = append(message.Reactions, models.Reaction{
			Name:  reaction.Name,
			Count: reaction.Count,
			Users: reaction.Users,
		})
	}

	// スレッド返信の処理
	if includeThreads && msg.ThreadTimestamp != "" && msg.ThreadTimestamp == msg.Timestamp {
		replies, err := c.getThreadReplies(channelID, msg.Timestamp)
		if err != nil {
			if c.verbose {
				fmt.Printf("警告: スレッド返信の取得に失敗しました: %s\n", err)
			}
		} else {
			message.ThreadReplies = replies
		}
	}

	return message, nil
}

// getThreadReplies はスレッドの返信を取得します
func (c *Client) getThreadReplies(channelID, threadTS string) ([]models.Message, error) {
	params := slack.GetConversationRepliesParameters{
		ChannelID: channelID,
		Timestamp: threadTS,
	}

	replies := []models.Message{}

	// スレッド返信の取得（ページネーションは簡略化）
	history, _, _, err := c.api.GetConversationReplies(&params)
	if err != nil {
		return nil, errors.Wrap(err, "スレッド返信の取得に失敗しました")
	}

	// 最初のメッセージはスレッドの親なのでスキップ
	for i, msg := range history {
		if i == 0 && msg.ThreadTimestamp == msg.Timestamp {
			continue
		}

		reply, err := c.convertToMessage(msg, channelID, false)
		if err != nil {
			if c.verbose {
				fmt.Printf("警告: 返信メッセージの変換に失敗しました: %s\n", err)
			}
			continue
		}
		replies = append(replies, *reply)
	}

	return replies, nil
}

// getUserName はユーザーIDからユーザー名を取得します
func (c *Client) getUserName(userID string) string {
	if userID == "" {
		return ""
	}

	user, err := c.api.GetUserInfo(userID)
	if err != nil {
		if c.verbose {
			fmt.Printf("警告: ユーザー情報の取得に失敗しました: %s\n", err)
		}
		return userID
	}
	return user.Name
}

// parseTimestamp はSlackのタイムスタンプをtime.Timeに変換します
func parseTimestamp(timestamp string) (time.Time, error) {
	// Slackのタイムスタンプは "1234567890.123456" の形式
	var sec float64

	_, err := fmt.Sscanf(timestamp, "%f", &sec)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "タイムスタンプのフォーマットが不正です")
	}

	// 整数部分のみを使用して時間を作成
	secInt := int64(sec)
	nsec := int64((sec - float64(secInt)) * 1e9)

	return time.Unix(secInt, nsec), nil
}

package api

import (
	"fmt"

	"github.com/slack-go/slack"
)

// Client はSlack APIとのやり取りを行うクライアントです
type Client struct {
	api *slack.Client
}

// NewClient は新しいSlack APIクライアントを作成します
func NewClient(token string) *Client {
	return &Client{
		api: slack.New(token),
	}
}

// GetChannelHistory はチャンネルの履歴を取得します
func (c *Client) GetChannelHistory(channelID string) ([]slack.Message, error) {
	params := slack.GetConversationHistoryParameters{
		ChannelID: channelID,
	}

	history, err := c.api.GetConversationHistory(&params)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel history: %w", err)
	}

	return history.Messages, nil
}

package logram

import "github.com/iluxaorlov/logram/internal/telegram"

type telelog struct {
	client *telegram.Client
	chatId int64
}

func NewWriter(token string, chatId int64) *telelog {
	client := telegram.NewClient(token)

	return &telelog{
		client: client,
		chatId: chatId,
	}
}

func (t *telelog) Write(p []byte) (n int, err error) {
	if err := t.client.Send(string(p), t.chatId); err != nil {
		return 0, err
	}

	return len(p), nil
}

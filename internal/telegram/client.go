package telegram

import (
	"bytes"
	"encoding/json"
	"github.com/iluxaorlov/logram/internal/model"
	"net/http"
)

const maxTextLength = 4096

type Client struct {
	url   string
	token string
}

func NewClient(token string) *Client {
	return &Client{
		url:   "https://api.telegram.org/bot",
		token: token,
	}
}

func (t *Client) Send(text string, chatId int64) error {
	var chunk []rune

	for _, char := range text {
		chunk = append(chunk, char)

		if len(chunk) >= maxTextLength {
			if err := t.SendMessage(string(chunk), chatId); err != nil {
				return err
			}

			chunk = nil
		}
	}

	return t.SendMessage(string(chunk), chatId)
}

func (t *Client) SendMessage(text string, chatId int64) error {
	message := &model.Message{
		ChatId: chatId,
		Text:   text,
	}

	var body bytes.Buffer

	if err := json.NewEncoder(&body).Encode(message); err != nil {
		return err
	}

	res, err := http.Post(t.url+t.token+"/sendMessage", "application/json", &body)
	if err != nil {
		return err
	}

	return res.Body.Close()
}

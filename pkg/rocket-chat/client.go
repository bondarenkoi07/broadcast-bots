package client

import (
	"fmt"

	"github.com/badkaktus/gorocket"
)

type Client interface {
	Login(login string, password string) error
	CreateMessage(login string, text string) gorocket.Message
	SendMessage(msg gorocket.Message) error
}

type client struct {
	client *gorocket.Client
}

func NewClient(host string) *client {
	c := gorocket.NewClient(host)
	return &client{client: c}
}

func (r client) Login(login string, password string) error {
	auth := gorocket.LoginPayload{
		User:     login,
		Password: password,
	}

	_, err := r.client.Login(&auth)

	return err
}

func (r client) CreateMessage(login string, text string) gorocket.Message {
	return gorocket.Message{
		Channel: login,
		Text:    text,
	}
}

func (r client) SendMessage(msg gorocket.Message) error {
	res, err := r.client.PostMessage(&msg)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("could not send message to %s, cause: %s", msg.Channel, res.Error)
	}

	return nil
}

package senders

import (
	"fmt"
	"log"
	"service-desk-senders/internal/domain"
	rc "service-desk-senders/pkg/rocket-chat"
)

type rocketChat struct {
	client rc.Client
	node   domain.Sender
}

func NewRocketChat(client rc.Client, node domain.Sender) *rocketChat {
	return &rocketChat{client: client, node: node}
}

func (s rocketChat) Send(message domain.Message) domain.Sender {
	for _, login := range message.UserRocketChatLogin {
		msg := s.client.CreateMessage(login, fmt.Sprintf("%v: %s", message.Event, message.Text))
		err := s.client.SendMessage(msg)
		if err != nil {
			log.Println(err)
		}
	}
	return s.node
}

func (s rocketChat) Set(sender domain.Sender) {
	s.node = sender
}

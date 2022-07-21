package domain

type Sender interface {
	Send(message Message) error
}

package domain

type Sender interface {
	Send(message Message) Sender
	Set(sender Sender)
}

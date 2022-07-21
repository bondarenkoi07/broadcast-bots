package domain

type Event string

const (
	EventMessage     Event = "message"
	EventCancelled   Event = "cancelled"
	EventReturned    Event = "returned"
	EventCompleted   Event = "completed"
	EventDeleted     Event = "deleted"
	EventChangeGroup Event = "change_group"
)

type Message struct {
	Event               Event    `json:"event"`
	Text                string   `json:"text"`
	UserRocketChatLogin []string `json:"user_rocket_chat_login"`
}

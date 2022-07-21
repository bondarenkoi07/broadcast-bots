package endpoints

import (
	"encoding/json"
	"net/http"
	"service-desk-senders/internal/domain"
)

type ServiceDeskEvent struct {
	root domain.Sender
}

func NewServiceDeskEvent(root domain.Sender) *ServiceDeskEvent {
	return &ServiceDeskEvent{root: root}
}

func (e ServiceDeskEvent) ReceiveEvent(w http.ResponseWriter, r *http.Request) {
	var msg domain.Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var node = e.root
	for node != nil {
		node = node.Send(msg)
	}
}

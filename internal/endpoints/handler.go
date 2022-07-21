package endpoints

import (
	"encoding/json"
	"net/http"
	"service-desk-bots/internal/domain"
)

type ServiceDeskEvent struct {
	root domain.Sender
}

func (e ServiceDeskEvent) ReceiveEvent(w http.ResponseWriter, r *http.Request) {
	var msg domain.Message

	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = e.root.Send(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

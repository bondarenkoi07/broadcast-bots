package main

import (
	"log"
	"net/http"
	"os"
	"service-desk-senders/internal/endpoints"
	"service-desk-senders/internal/senders"
	client "service-desk-senders/pkg/rocket-chat"
	"time"
)

func main() {
	rocketChatLogin := os.Getenv("ROCKET_CHAT_BOT_LOGIN")
	rocketChatPassword := os.Getenv("ROCKET_CHAT_BOT_PASSWORD")
	rocketChatHost := os.Getenv("ROCKET_CHAT_BOT_HOST")
	host := os.Getenv("SERVICE_DESK_BOTS_HOST")

	cl := client.NewClient(rocketChatHost)
	err := cl.Login(rocketChatLogin, rocketChatPassword)
	if err != nil {
		log.Fatal(err)
	}

	sender := senders.NewRocketChat(cl, nil)

	ep := endpoints.NewServiceDeskEvent(sender)

	mux := http.NewServeMux()
	mux.HandleFunc("api/event", ep.ReceiveEvent)

	srv := &http.Server{
		Handler: mux,
		Addr:    host,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 55 * time.Second,
		ReadTimeout:  55 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

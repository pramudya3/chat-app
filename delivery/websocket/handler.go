package websocket

import (
	"chat-app/domain"
	"log"
	"net/http"
)

type handlerWebsocket struct {
	manager *domain.Manager
}

func NewHandlerWebsocket(manager *domain.Manager) *handlerWebsocket {
	return &handlerWebsocket{
		manager: manager,
	}
}

func (h *handlerWebsocket) ServeWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := domain.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrade websocket: %v", err)
		return
	}

	client := &domain.Client{
		Manager: h.manager,
		Conn:    conn,
		Send:    make(chan []byte, 256),
	}

	client.Manager.Register <- client

	go client.Write()
	go client.Read()
}

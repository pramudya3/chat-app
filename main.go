package main

import (
	ws "chat-app/delivery/websocket"
	"chat-app/domain"
	"log"
	"net/http"
)

func main() {
	manager := domain.NewManager()
	go manager.Run()

	handler := ws.NewHandlerWebsocket(manager)

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/ws", handler.ServeWebsocket)

	log.Println("server is starting at :8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

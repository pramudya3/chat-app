package domain

import (
	"encoding/json"
	"log"
)

type Manager struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}

func NewManager() *Manager {
	return &Manager{
		Clients:    map[*Client]bool{},
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (m *Manager) Run() {
	for {
		select {
		case client := <-m.Register:
			m.Clients[client] = true
			msgByte, err := json.Marshal(&Message{Type: Register, Sender: client.Username})
			if err != nil {
				log.Printf("error marshaling json, %v", err)
				return
			}
			m.send(msgByte, client)
		case client := <-m.Unregister:
			msgByte, err := json.Marshal(&Message{Type: Unregister, Sender: client.Username})
			if err != nil {
				log.Printf("error marshaling json, %v", err)
				return
			}
			m.send(msgByte, client)

			close(client.Send)
			delete(m.Clients, client)

		case msg := <-m.Broadcast:
			for client := range m.Clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(m.Clients, client)
				}
			}
		}
	}
}

func (m *Manager) send(msg []byte, ignore *Client) {
	for conn := range m.Clients {
		if conn != ignore {
			conn.Send <- msg
		}
	}
}

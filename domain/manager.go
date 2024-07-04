package domain

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
		case client := <-m.Unregister:
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

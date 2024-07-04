package domain

const (
	Register   = "Register"
	Unregister = "Unregister"
	Chat       = "Chat"
)

type Message struct {
	Sender  string `json:"sender,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

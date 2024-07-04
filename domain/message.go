package domain

type Message struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

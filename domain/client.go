package domain

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 5 * time.Second
	pongWait       = 20 * time.Second
	pingPeriod     = 10 * time.Second
	maxMessageSize = 512
	bufferSize     = 1024
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  bufferSize,
	WriteBufferSize: bufferSize,
}

type Client struct {
	Manager *Manager

	*websocket.Conn

	Send chan []byte

	Username string
}

func (c *Client) Read() {
	defer func() {
		c.Manager.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msgByte, err := json.Marshal(&Message{Sender: c.Username, Message: string(msg)})
		if err != nil {
			log.Printf("error marshal JSON, %v", err)
			return
		}

		c.Manager.Broadcast <- msgByte
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, msg)

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("error ping websocket server: %v", err)
				return
			}
		}
	}
}

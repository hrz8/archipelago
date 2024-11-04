package chat

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Chat struct {
	roomCapacity  int
	wsConnections map[string]*websocket.Conn
	usersLock     sync.Mutex
}

func NewChatService(cap int) *Chat {
	return &Chat{
		roomCapacity:  cap,
		wsConnections: make(map[string]*websocket.Conn),
	}
}

func (c *Chat) handleConnected() {
	for userKey, userConn := range c.wsConnections {
		var otherKey string
		if userKey == "Person A" {
			otherKey = "B"
		} else {
			otherKey = "A"
		}
		userConn.WriteMessage(websocket.TextMessage, []byte("Chat started! You are now connected with "+otherKey))
	}
}

func (c *Chat) handleMessages(person string, conn *websocket.Conn) {
	// this defer function will be invoked when there is close signal triggered
	defer func() {
		c.usersLock.Lock()
		defer c.usersLock.Unlock()

		for _, userConn := range c.wsConnections {
			userConn.WriteMessage(websocket.TextMessage, []byte("User disconnected, please refresh."))
			userConn.Close()
		}
		// reset websocket connection
		c.wsConnections = make(map[string]*websocket.Conn)
	}()

	// wait incoming message
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return
		}

		for otherPerson, otherConn := range c.wsConnections {
			if otherPerson != person {
				otherConn.WriteMessage(messageType, message)
			}
		}
	}
}

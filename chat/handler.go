package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	chatService *Chat
}

func NewHandler() *Handler {
	return &Handler{
		chatService: NewChatService(2),
	}
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	h.chatService.usersLock.Lock()
	defer h.chatService.usersLock.Unlock()

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "ws failed", http.StatusInternalServerError)
		return
	}

	if len(h.chatService.wsConnections) >= h.chatService.roomCapacity {
		conn.WriteMessage(websocket.TextMessage, []byte("Room is full!"))
		conn.Close()
		return
	}

	var person string
	if len(h.chatService.wsConnections) == 0 {
		person = "Person A"
	} else {
		person = "Person B"
	}
	h.chatService.wsConnections[person] = conn
	conn.WriteMessage(websocket.TextMessage, []byte("Connected! You are "+person))

	if len(h.chatService.wsConnections) == h.chatService.roomCapacity {
		h.chatService.handleConnected()
	}

	go h.chatService.handleMessages(person, conn)
}

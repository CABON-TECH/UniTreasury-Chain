package ws
import (
	"encoding/json"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}
type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
	log     *zap.Logger
}
func NewHub(log *zap.Logger) *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
		log:     log,
	}
}
func (h *Hub) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.log.Error("websocket upgrade failed", zap.Error(err))
		return
	}
	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
	h.log.Info("websocket client connected", zap.String("remote", r.RemoteAddr))
	go func() {
		defer func() {
			h.mu.Lock()
			delete(h.clients, conn)
			h.mu.Unlock()
			conn.Close()
			h.log.Info("websocket client disconnected", zap.String("remote", r.RemoteAddr))
		}()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}()
}
func (h *Hub) Broadcast(message interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()
	msgBytes, err := json.Marshal(message)
	if err != nil {
		h.log.Error("websocket broadcast marshal error", zap.Error(err))
		return
	}
	for client := range h.clients {
		err := client.WriteMessage(websocket.TextMessage, msgBytes)
		if err != nil {
			client.Close()
			delete(h.clients, client)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
)

type Client struct {
	Username string `json:"username"`
	Conn     *websocket.Conn
}

type Registry struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]*Client
}

var Connects = Registry{
	clients: make(map[*websocket.Conn]*Client),
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Printf("Error connecting: %s", err)
		return
	}

	Connects.mu.Lock()
	Connects.clients[conn] = &Client{
		Username: username,
		Conn:     conn,
	}
	Connects.mu.Unlock()

	defer conn.Close(websocket.StatusInternalError, "Unexpected close")
	ctx := r.Context()
	for {

		msgType, msg, err := conn.Read(ctx)
		if err != nil {
			Connects.mu.Lock()
			delete(Connects.clients, conn)
			Connects.mu.Unlock()
			conn.Close(websocket.StatusNormalClosure, "bye")
			return
		}

		Connects.mu.Lock()
		for _, client := range Connects.clients {
			if client.Conn == conn {
				continue
			}
			if err := client.Conn.Write(ctx, msgType, fmt.Appendf([]byte(""), "%s: %s", username, msg)); err != nil {
				fmt.Printf("Error Writing Message: %s", err)
				break
			}
		}
		Connects.mu.Unlock()
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/ws", HandleWebSocket)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8001),
		Handler: mux,
	}

	fmt.Println("Server running on port: 8001")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s", err)
		return
	}
}

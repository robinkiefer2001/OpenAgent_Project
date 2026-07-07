package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupRoutes() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/ws", tunnel)
}

func connectTunnel(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(p))
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server is Reachable")
	fmt.Println("Endpoint Hit: status requested")
}

func tunnel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: tunnel requested")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "websocket upgrade failed", http.StatusBadRequest)
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	connectTunnel(ws)
}

func main() {
	fmt.Println("Brocker Server is starting on port 8080...")
	setupRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Broker server failed to start:", err)
	}
}

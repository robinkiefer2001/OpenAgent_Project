package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func connectToBroker() {
	resp, err := http.Get("http://localhost:8080/status")
	if err != nil {
		fmt.Println("Error connecting to broker:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Broker is reachable. Status:", resp.Status)
	} else {
		fmt.Println("Failed to reach broker. Status:", resp.Status)
	}
}

func openTunnel() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		fmt.Println("Error connecting to broker tunnel:", err)
		return
	}
	defer conn.Close()

	message := "Hello from client!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	_, p, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}
	fmt.Println("Received from broker:", string(p))
}

func main() {
	connectToBroker()
	openTunnel()
}

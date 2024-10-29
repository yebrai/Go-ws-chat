package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// Configuration for the WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow connections from any origin
}

// Structure for the message
type Message struct {
	Name    string `json:"name"`    // The name of the user
	Message string `json:"message"` // The content of the message
}

// Handle WebSocket connections
func handleConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close() // Ensure the connection is closed at the end

	fmt.Println("New client connected")

	for {
		var msg Message
		// Read the message in JSON format
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Message received from %s: %s\n", msg.Name, msg.Message)

		// Send the message back to the client
		err = conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}

func main() {
	// Handle the route for WebSocket connections
	http.HandleFunc("/ws", handleConnection)

	// Serve static files (index.html)
	http.Handle("/", http.FileServer(http.Dir("./")))

	fmt.Println("Server listening on :3000")
	err := http.ListenAndServe(":3000", nil) // Start the server on port 3000
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

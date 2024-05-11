package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	clients     = make(map[string]struct{})
	clientsLock sync.Mutex
)

func main() {
	http.HandleFunc("/numbers", numberHandler)
	http.HandleFunc("/connected-clients", connectedClientsHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	clientID := fmt.Sprintf("%d", time.Now().UnixNano()) // timestamp as identifier
	defer func() {
		// Remove the client when connection is closed
		clientsLock.Lock()
		delete(clients, clientID)
		clientsLock.Unlock()
	}()

	// Add the client to the list
	clientsLock.Lock()
	clients[clientID] = struct{}{}
	clientsLock.Unlock()

	// Send an incrementing number to the client every second
	number := 0
	for {
		select {
		case <-r.Context().Done():
			// Connection closed, stop sending data
			return
		default:
			fmt.Fprintf(w, "data: %d\n\n", number)
			w.(http.Flusher).Flush()

			number++
			time.Sleep(1 * time.Second)
		}
	}
}

func connectedClientsHandler(w http.ResponseWriter, r *http.Request) {
	// Convert map keys to slice
	clientsLock.Lock()
	defer clientsLock.Unlock()

	clientList := make([]string, 0, len(clients))
	for client := range clients {
		clientList = append(clientList, client)
	}

	// Marshal the client list to JSON
	clientJSON, err := json.Marshal(clientList)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(clientJSON)
}

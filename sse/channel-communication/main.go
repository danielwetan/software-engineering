package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	messageChan = make(chan eventStream)
	closeChan = make(chan struct{})

	http.HandleFunc("/messages", getMessageHandler)
	http.HandleFunc("/send", sendMessageHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type eventStream struct {
	ID   string
	Type string
	Data string
}

var messageChan chan eventStream
var closeChan chan struct{}

func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	for {
		select {
		case msg := <-messageChan:
			fmt.Fprintf(w, "id: %s\n", msg.ID)       // event ID
			fmt.Fprintf(w, "event: %s\n", msg.Type)  // event type
			fmt.Fprintf(w, "data: %s\n\n", msg.Data) // we can also send event in json string format
			flusher.Flush()                          // Flush the buffer to send immediately
		case <-closeChan:
			return
		}
	}
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	msgID := fmt.Sprintf("%d", time.Now().UnixNano()) // timestamp as identifier
	eventType := r.URL.Query().Get("type")
	data := r.URL.Query().Get("data")
	messageChan <- eventStream{
		ID:   msgID,
		Type: eventType,
		Data: data,
	}

	if eventType == "close" {
		closeChan <- struct{}{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

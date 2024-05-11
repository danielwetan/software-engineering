package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/numbers", numberHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	// send a increment number every second
	number := 0
	for {
		fmt.Fprintf(w, "data: %d \n\n", number)
		w.(http.Flusher).Flush()

		number += 1
		time.Sleep(1 * time.Second)
	}
}

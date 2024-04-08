package handlers

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	ID int `json:"id,string"`
	Content string `json:"content"`
}

var messages = make([]Message, 0)

func appendMessage(message Message) {
	messages = append(messages, message)
}

func MessageEndpoint(w http.ResponseWriter, r *http.Request) {
	switch (r.Method) {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(messages)
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		var message Message

		jsonerr := json.NewDecoder(r.Body).Decode(&message)
		defer r.Body.Close()
		if (jsonerr != nil) {
			http.Error(w, jsonerr.Error(), http.StatusBadRequest)
			return
		}

		appendMessage(message)

		json.NewEncoder(w).Encode(message)

	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

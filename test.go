package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}
type Feedback struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var message Message
		var feedback Feedback

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&message)
		if err != nil {
			feedback.Status = "400"
			feedback.Message = "Error decoding JSON"

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(feedback)
			return
		}
		if message.Message == ""{
			feedback.Status = "400"
			feedback.Message = "Отсутствует поле 'message' в JSON"

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(feedback)
			return
		}
		fmt.Printf("Received JSON data: %+v\n", message.Message)

		feedback.Status = "Success"
		feedback.Message = "Данные успешно приняты"
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(feedback)
	})

	http.ListenAndServe(":8080", nil)
}

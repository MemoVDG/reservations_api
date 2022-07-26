package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error parsing data")
	}

}

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func RespondJSON(w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("encode response body to json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		rsp := &Response{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			log.Fatalf("encode response body to json: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		log.Fatalf("encode response body to json: %v", err)
	}
}

package types

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Data       any `json:"data"`
	statusCode int
}

func NewJsonResponse(Data any, statusCode int) *JsonResponse {
	return &JsonResponse{Data: Data, statusCode: statusCode}
}

func (jr *JsonResponse) Respond(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(jr.statusCode)
	err := json.NewEncoder(w).Encode(jr.Data)
	if err != nil {
		log.Println("Error, while sending data to client")
	}
}

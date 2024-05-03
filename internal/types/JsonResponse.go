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
	/*w.Header().Set("Access-Control-Allow-Origin", "*")*/
	/*w.Header().Set("Access-Control-Allow-Methods", "GET, POST, HEAD, OPTIONS")*/
	/*w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")*/

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(jr.statusCode)
	err := json.NewEncoder(w).Encode(jr.Data)
	if err != nil {
		log.Println("Error, while sending data to client")
	}
}

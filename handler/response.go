package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-product-api/models"
)

func GenerateError(rw http.ResponseWriter, message string, status_code int ){
	WriteJSON(rw, status_code, &models.ErrorResponse{
		Message: message,
		StatusCode: status_code,
	})
}

func WriteJSON(rw http.ResponseWriter, status_code int, data any){
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status_code)
	if err := json.NewEncoder(rw).Encode(data); err != nil{
		log.Println("Encoding response error: ", err)
		return
	}
}
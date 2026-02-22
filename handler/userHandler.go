package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-product-api/models"
	"simple-product-api/service"
	"strconv"
	"strings"
)

type UserHandler struct {
	Service *service.UserService
}

func (uh *UserHandler) GetAllUsers(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type", "application/json")

	data, err := uh.Service.GetAllUsers()
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(data)
	if err != nil {
		//server-side error
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
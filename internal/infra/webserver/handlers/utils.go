package handlers

import (
	"encoding/json"
	"net/http"
)

func handlerBadResquest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	msg := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(msg)
}

package _lib

import (
	"encoding/json"
	"net/http"
)

func MarshalJsonAndWrite(v interface{}, w http.ResponseWriter) {
	payload, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}

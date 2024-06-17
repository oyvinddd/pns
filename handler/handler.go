package handler

import (
	"encoding/json"
	"net/http"
)

func RegisterDeviceToken(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		var tokenRequest = struct {
			Token string `json:"token"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&tokenRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DeleteDeviceToken(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

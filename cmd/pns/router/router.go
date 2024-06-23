package router

import (
	"encoding/json"
	"net/http"
)

const (
	tokensEndpoint  = "/api/v1/tokens"
	messageEndpoint = "/api/v1/messages"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(tokensEndpoint, handleDeviceTokens)
	mux.HandleFunc(messageEndpoint, sendPushMessage)
	return mux
}

func handleDeviceTokens(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		registerDeviceToken(w, r)
	case "DELETE":
		deleteDeviceToken(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func sendPushMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Register device token
func registerDeviceToken(w http.ResponseWriter, r *http.Request) {
	var tokenRequest = struct {
		Token string `json:"token"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&tokenRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete device token
func deleteDeviceToken(w http.ResponseWriter, r *http.Request) {
}

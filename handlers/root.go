package handlers

import (
	"net/http"

	"github.com/swift-sunshine/swscore/status"
)

func Root(w http.ResponseWriter, r *http.Request) {
	getStatus(w, r)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	RespondWithJSONIndent(w, http.StatusOK, status.Get())
}

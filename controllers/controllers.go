package controllers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"version": "1.0",
	}
	json.NewEncoder(w).Encode(response)
}

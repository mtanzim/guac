package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type LoginArg struct {
	Username string
	Password string
}

func LoginController(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	expectedUser, expectedPass := os.Getenv("USERNAME"), os.Getenv("USERPASS")
	if expectedUser == "" || expectedPass == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("User/password env not setup")
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t LoginArg
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(expectedUser, expectedPass)
	log.Println(t)
	if t.Username == expectedUser && t.Password == expectedPass {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Token")
		return
	}
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode("Unauthorized")
}

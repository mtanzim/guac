package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mtanzim/guac/server/auth"

	"github.com/dgrijalva/jwt-go"
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

	expectedUser, expectedPass, secret := os.Getenv("USERNAME"), os.Getenv("USERPASS"), os.Getenv("SECRET")
	if expectedUser == "" || expectedPass == "" || secret == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("user/password/secret env not setup")
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t LoginArg
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if t.Username == expectedUser && t.Password == expectedPass {
		claims := auth.MyCustomClaims{
			t.Username,
			jwt.StandardClaims{
				// in ms
				ExpiresAt: (time.Now().Unix() * 1000) + time.Hour.Milliseconds(),
				Issuer:    "guac",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Failed to sign token")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tokenString)
		return
	}
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode("Unauthorized")
}

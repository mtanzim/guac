package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// TODO: verify auth, and return a token
func AuthVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		expectedUser, secret := os.Getenv("USERNAME"), os.Getenv("SECRET")
		if expectedUser == "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("User/password env not setup")
			return
		}

		reqToken := req.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Malformed token")
			return
		}
		parsedToken := strings.TrimSpace(splitToken[1])
		token, err := jwt.ParseWithClaims(parsedToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			log.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
			if claims.Username == expectedUser && claims.ExpiresAt > (time.Now().Unix()*1000) {
				next.ServeHTTP(w, req)
				return
			}
		}
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Unauthorized")

	})
}

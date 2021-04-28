package auth

import (
	"encoding/json"
	"net/http"
	"os"
)

// TODO: verify auth, and return a token
func AuthVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		expectedUser, expectedPass := os.Getenv("USERNAME"), os.Getenv("USERPASS")
		if expectedUser == "" || expectedPass == "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("User/password env not setup")
			return
		}

		// TODO: do not pass passwords in the query params!
		user := req.URL.Query().Get("user")
		pass := req.URL.Query().Get("pass")

		if user == expectedUser && pass == expectedPass {
			next.ServeHTTP(w, req)
			return
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Unauthorized")

	})
}

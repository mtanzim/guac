package controllers

import (
	"fmt"
	"net/http"
)

func HealthController(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	fmt.Fprintf(w, "OK")
}

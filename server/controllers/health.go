package controllers

import (
	"fmt"
	"net/http"
)

func HealthController(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
}

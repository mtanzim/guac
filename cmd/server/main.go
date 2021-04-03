package main

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/server"
)

func main() {
	http.HandleFunc("/hello", server.Hello)
	http.ListenAndServe(":8090", nil)
}

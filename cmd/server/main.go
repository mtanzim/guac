package main

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/server"
)

func main() {
	http.HandleFunc("/data", server.Data)
	http.ListenAndServe(":8090", nil)
}

package main

import (
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/server"
)

func main() {
	http.HandleFunc("/data", server.DataController)
	http.HandleFunc("/plot", server.PlotController)
	http.ListenAndServe(":"+os.Getenv("REST_PORT"), nil)
}

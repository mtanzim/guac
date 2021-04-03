package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/server/controllers"
)

func main() {
	http.HandleFunc("/data", controllers.DataController)
	http.HandleFunc("/plot", controllers.PlotController)
	port := os.Getenv("REST_PORT")
	if port == "" {
		log.Fatalln("Please provide env variable for REST_PORT")
	}
	log.Println("Starting server on PORT:" + port)
	http.ListenAndServe(":"+port, nil)
}

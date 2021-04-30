package server

import (
	"log"
	"net/http"
	"os"

	"github.com/mtanzim/guac/server/auth"
	"github.com/mtanzim/guac/server/controllers"
)

func Start() {
	http.HandleFunc("/health", controllers.HealthController)
	http.HandleFunc("/login", controllers.LoginController)
	http.Handle("/", auth.AuthVerify(http.HandlerFunc(controllers.RootController)))
	http.Handle("/data", auth.AuthVerify(http.HandlerFunc(controllers.DataController)))
	http.Handle("/plot", auth.AuthVerify(http.HandlerFunc(controllers.PlotController)))

	port := os.Getenv("REST_PORT")
	if port == "" {
		log.Fatalln("Please provide env variable for REST_PORT")
	}
	log.Println("Starting server on PORT:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

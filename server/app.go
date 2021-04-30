package server

import (
	"log"
	"net/http"
	"os"

	"github.com/mtanzim/guac/server/auth"
	"github.com/mtanzim/guac/server/controllers"
)

var (
	ApiURL = "/api/v1"
)

func Start() {
	http.HandleFunc(ApiURL+"/health", controllers.HealthController)
	http.HandleFunc(ApiURL+"/login", controllers.LoginController)
	http.Handle(ApiURL+"/", auth.AuthVerify(http.HandlerFunc(controllers.RootController)))
	http.Handle(ApiURL+"/data", auth.AuthVerify(http.HandlerFunc(controllers.DataController)))
	http.Handle(ApiURL+"/plot", auth.AuthVerify(http.HandlerFunc(controllers.PlotController)))
	http.Handle("/", http.FileServer(http.Dir("./public")))

	port := os.Getenv("REST_PORT")
	if port == "" {
		log.Fatalln("Please provide env variable for REST_PORT")
	}
	log.Println("Starting server on PORT:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

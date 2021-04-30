package server

import (
	"log"
	"net/http"
	"os"

	"github.com/mtanzim/guac/server/auth"
	"github.com/mtanzim/guac/server/controllers"
)

var (
	BaseURL = "/api/v1"
)

func Start() {

	http.HandleFunc(BaseURL+"/health", controllers.HealthController)
	http.HandleFunc(BaseURL+"/login", controllers.LoginController)
	http.Handle(BaseURL+"/", auth.AuthVerify(http.HandlerFunc(controllers.RootController)))
	http.Handle(BaseURL+"/data", auth.AuthVerify(http.HandlerFunc(controllers.DataController)))
	http.Handle(BaseURL+"/plot", auth.AuthVerify(http.HandlerFunc(controllers.PlotController)))

	port := os.Getenv("REST_PORT")
	if port == "" {
		log.Fatalln("Please provide env variable for REST_PORT")
	}
	log.Println("Starting server on PORT:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

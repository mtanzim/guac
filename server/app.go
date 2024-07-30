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

func allowCORS(next http.Handler) http.Handler {
	if os.Getenv("IS_DEV") != "true" {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(w, req)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, expect")
		if req.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, req)
	})
}

func Start() {

	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("./public")))
	router.HandleFunc(ApiURL+"/health", controllers.HealthController)
	router.Handle(ApiURL+"/login", allowCORS(http.HandlerFunc(controllers.LoginController)))
	router.Handle(ApiURL+"/data", allowCORS(auth.AuthVerify(http.HandlerFunc(controllers.DataController))))
	// Backend plots are disabled
	// http.Handle(ApiURL+"/", auth.AuthVerify(http.HandlerFunc(controllers.RootController)))
	// http.Handle(ApiURL+"/plot", http.HandlerFunc(controllers.PlotController))

	port := os.Getenv("REST_PORT")
	if port == "" {
		log.Fatalln("Please provide env variable for REST_PORT")
	}
	log.Println("Starting server on PORT:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

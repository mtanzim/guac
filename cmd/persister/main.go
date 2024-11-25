package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/firestoreClient"
	"github.com/mtanzim/guac/wakaApi"
)

func main() {
	data := wakaApi.TransformData()
	log.Println("successfully got api data")
	firestoreClient.PutData(data)
	log.Println("successfully placed data on firestore")
}

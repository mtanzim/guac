package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/firestoreClient"
	"github.com/mtanzim/guac/wakaApi"
)

func main() {
	data := wakaApi.TransformData()
	firestoreClient.PutData(data)
}

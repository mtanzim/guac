package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/firestore"
	"github.com/mtanzim/guac/wakaApi"
)

func main() {
	data := wakaApi.TransformData()
	firestore.PutData(data)
}

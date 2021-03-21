package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/wakaApi"
)

func main() {
	data := wakaApi.TransformData()
	log.Println(data)

}

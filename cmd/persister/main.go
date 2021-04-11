package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/wakaApi"
)

func main() {
	data := wakaApi.TransformData()
	dynamo.PutData(data)
}

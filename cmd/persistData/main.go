package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/wakaApi"
)

type Item struct {
	Date string
	Data interface{}
}

func main() {
	data := wakaApi.TransformData()
	dynamo.PutData(data)
}

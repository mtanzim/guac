package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/firestore"
)

func main() {
	firestore.GetData("2020-06-01", "2021-06-03")
}

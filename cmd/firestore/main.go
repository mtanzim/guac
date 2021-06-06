package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/firestore"
)

func main() {
	firestore.MigrateDynamo()
}

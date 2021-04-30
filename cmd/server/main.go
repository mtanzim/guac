package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/server"
)

func main() {
	server.Start()
}

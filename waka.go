package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Hello")
	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	sEnc := b64.StdEncoding.EncodeToString([]byte(apiKey))
	authHeader := "Basic " + sEnc

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", authHeader)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var v interface{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("%s", body)
	log.Printf("%s", v)
}

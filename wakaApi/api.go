package wakaApi

import (
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mtanzim/guac/utils"
)

func getAuthHeader() string {
	apiKey := os.Getenv("API_KEY")
	sEnc := b64.StdEncoding.EncodeToString([]byte(apiKey))
	authHeader := "Basic " + sEnc
	return authHeader
}

func getData() []byte {
	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	authHeader := getAuthHeader()
	req.Header.Add("Authorization", authHeader)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	return body
}

func TransformData() map[string]interface{} {
	body := getData()

	var v map[string]interface{}
	err := json.Unmarshal(body, &v)
	if err != nil {
		log.Fatal(err)
	}

	if v["data"] == nil {
		utils.PrettyPrint(v)
		log.Panicln(v["error"])
	}
	data := v["data"].([]interface{})
	transformedData := make(map[string]interface{})
	for _, d := range data {
		var date string
		switch dd := d.(type) {
		case map[string]interface{}:
			rng := dd["range"].(map[string]interface{})
			date = rng["date"].(string)
			transformedData[date] = d
		default:
			// Do nothing
		}

	}
	return transformedData
}

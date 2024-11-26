package main

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/mtanzim/guac/firestoreClient"
	"github.com/mtanzim/guac/wakaApi"

	_ "github.com/joho/godotenv/autoload"
)

func jobFn() {
	data := wakaApi.TransformData()
	log.Println("successfully got api data")
	firestoreClient.PutData(data)
	log.Println("successfully placed data on firestore")
}

type Config struct {
	apiKey    string
	baseUrl   string
	collName  string
	projectID string

	// Job-defined
	taskNum    string
	attemptNum string
}

func configFromEnv() (Config, error) {
	// Job-defined
	taskNum := os.Getenv("CLOUD_RUN_TASK_INDEX")
	attemptNum := os.Getenv("CLOUD_RUN_TASK_ATTEMPT")

	// validate env vars
	apiKey := os.Getenv("API_KEY")
	baseUrl := os.Getenv("BASE_URL")
	collName := os.Getenv("GOOGLE_WAKA_COLL")
	projectID := os.Getenv("GOOGLE_PROJECT_ID")

	if apiKey == "" || baseUrl == "" || collName == "" || projectID == "" {
		return Config{}, errors.New("env vars not correctly configure")
	}

	config := Config{
		taskNum:    taskNum,
		attemptNum: attemptNum,
		apiKey:     apiKey,
		baseUrl:    baseUrl,
		collName:   collName,
		projectID:  projectID,
	}
	return config, nil
}

func sleepMsToInt(s string) (int64, error) {
	sleepMs, err := strconv.ParseInt(s, 10, 64)
	return sleepMs, err
}

func main() {
	config, err := configFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting Task #%s, Attempt #%s ...", config.taskNum, config.attemptNum)
	jobFn()
	log.Printf("Completed Task #%s, Attempt #%s", config.taskNum, config.attemptNum)
}

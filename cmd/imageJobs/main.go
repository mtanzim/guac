package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-echarts/snapshot-chromedp/render"
	_ "github.com/joho/godotenv/autoload"
	bucketClient "github.com/mtanzim/guac/gcpBucketClient"
	"github.com/mtanzim/guac/plotData"
	"github.com/mtanzim/guac/server/services"
)

const filePath = "guac-pie-plot.png"

type dateRange struct {
	start string
	end   string
}

func getDateRange(months int) dateRange {
	// Get the current date and time
	now := time.Now()

	// Create a new Date object with the year, month, and day from the current date
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	startDate := endDate.AddDate(0, -months, 0)

	// Format the date as a string in the format "2023-11-28"
	formattedEndDate := endDate.Format("2006-01-02")
	formattedStartDate := startDate.Format("2006-01-02")

	// Print the formatted date
	return dateRange{
		start: formattedStartDate,
		end:   formattedEndDate,
	}

}

func jobFn(months int, topK int, bucketName string) error {
	dr := getDateRange(months)
	log.Printf("plotting data for top %d languages with range %s to %s", topK, dr.start, dr.end)
	rv := services.DataService(dr.start, dr.end)
	pie := plotData.LanguagePieMinimal(rv.LangStats, rv.StartDate, rv.EndDate, int64(topK))
	err := render.MakeChartSnapshot(pie.RenderContent(), filePath)
	if err != nil {
		return err
	}
	bucketClient.UploadFile(bucketName, filePath, filePath)
	return nil
}

type Config struct {
	collName   string
	projectID  string
	months     int
	topK       int
	bucketName string
	// Job-defined
	taskNum    string
	attemptNum string
}

func configFromEnv() (Config, error) {
	// Job-defined
	taskNum := os.Getenv("CLOUD_RUN_TASK_INDEX")
	attemptNum := os.Getenv("CLOUD_RUN_TASK_ATTEMPT")

	// validate env vars
	collName := os.Getenv("GOOGLE_WAKA_COLL")
	bucketName := os.Getenv("GCP_BUCKET_NAME")
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	months := os.Getenv("PLOT_MONTHS")
	topK := os.Getenv("PLOT_TOPK")

	_, err := os.Open("./public/v1/colors.json")
	if err == nil {
		log.Println("found colors file")
	} else {
		return Config{}, err
	}

	if collName == "" || projectID == "" || months == "" || topK == "" || bucketName == "" {
		return Config{}, errors.New("env vars not correctly configured")
	}

	monthsN, err := strconv.Atoi(months)
	if err != nil {
		return Config{}, err
	}

	topKn, err := strconv.Atoi(topK)
	if err != nil {
		return Config{}, err
	}

	config := Config{
		taskNum:    taskNum,
		attemptNum: attemptNum,
		months:     monthsN,
		topK:       topKn,
		collName:   collName,
		projectID:  projectID,
		bucketName: bucketName,
	}
	return config, nil
}

func main() {
	config, err := configFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting Task #%s, Attempt #%s ...", config.taskNum, config.attemptNum)
	err = jobFn(config.months, config.topK, config.bucketName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Completed Task #%s, Attempt #%s", config.taskNum, config.attemptNum)
}

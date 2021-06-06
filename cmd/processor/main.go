package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/firestore"
	"github.com/mtanzim/guac/processData"
	"github.com/mtanzim/guac/utils"
)

type Item struct {
	Date string
	Data interface{}
}

func main() {
	start, end := os.Getenv("START"), os.Getenv("END")
	if start == "" || end == "" {
		log.Fatalln("Please specify start and end dates in .env")
	}
	// data := dynamo.GetData(start, end)
	data := firestore.Demo()
	dailyStats := processData.DailyTotal(data)
	actualStart, actualEnd := processData.GetDateRange(dailyStats)

	langStats := processData.LanguageSummary(data)
	projStats := processData.ProjectSummary(data)

	utils.PrettyPrint(dailyStats)
	utils.PrettyPrint(langStats.Durations)
	utils.PrettyPrint(langStats.Percentages)
	utils.PrettyPrint(projStats)

	log.Println(actualStart, actualEnd)

}

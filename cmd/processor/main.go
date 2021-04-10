package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/dynamo"
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
		log.Panicln("Please specify start and end dates in .env")
	}
	data := dynamo.GetData(start, end)
	dailyStats := processData.DailyTotal(data)
	actualStart, actualEnd := processData.GetDateRange(dailyStats)

	langStats := processData.LanguageSummary(data)
	utils.PrettyPrint(dailyStats)
	utils.PrettyPrint(langStats.Durations)
	utils.PrettyPrint(langStats.Percentages)
	log.Println(actualStart, actualEnd)

}

package main

import (
	"log"

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
	start, end := "2020-03-22", "2022-03-25"
	data := dynamo.GetData(start, end)
	dailyStats := processData.DailyTotal(data)
	actualStart, actualEnd := processData.GetDateRange(dailyStats)

	langStats := processData.LanguageSummary(data)
	utils.PrettyPrint(dailyStats)
	utils.PrettyPrint(langStats.Durations)
	utils.PrettyPrint(langStats.Percentages)
	log.Println(actualStart, actualEnd)

}

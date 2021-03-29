package main

import (
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
	// TODO: get from DynamoDB
	// data := wakaApi.TransformData()
	data := dynamo.GetData("2021-03-22", "2021-03-25")
	dailyStats := processData.DailyTotal(data)
	utils.PrettyPrint(dailyStats)
	// start, end := processData.GetDateRange(dailyStats)

	// langStats := processData.LanguageSummary(data)
	// log.Println(start, end)
	// utils.PrettyPrint(dailyStats)
	// utils.PrettyPrint(langStats.Durations)
	// utils.PrettyPrint(langStats.Percentages)

}

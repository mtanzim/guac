package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/plotData"
	"github.com/mtanzim/guac/processData"
	"github.com/mtanzim/guac/wakaApi"
)

type Item struct {
	Date string
	Data interface{}
}

func main() {
	// TODO: get from DynamoDB
	data := wakaApi.TransformData()
	dailyStats := processData.DailyTotal(data)
	start, end := processData.GetDateRange(dailyStats)
	langStats := processData.LanguageSummary(data)
	cleanLangPct := processData.CleanLangPct(langStats.Percentages)
	log.Println(start, end)
	plotData.DailyBarChart(dailyStats)
	plotData.LanguagePie(cleanLangPct)

}

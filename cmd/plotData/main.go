package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/plotData"
	"github.com/mtanzim/guac/processData"
)

type Item struct {
	Date string
	Data interface{}
}

func main() {
	start, end := "2020-01-01", "2022-12-31"
	data := dynamo.GetData(start, end)
	dailyStats := processData.DailyTotal(data)
	actualStart, actualEnd := processData.GetDateRange(dailyStats)
	langStats := processData.LanguageSummary(data)
	cleanLangPct := processData.CleanLangPct(langStats.Percentages)
	log.Println(actualStart, actualEnd)
	plotData.DailyBarChart(dailyStats, actualStart, actualEnd)
	plotData.LanguagePie(cleanLangPct, actualStart, actualEnd)

}

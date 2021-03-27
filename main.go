package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/processData"
	"github.com/mtanzim/guac/utils"
	"github.com/mtanzim/guac/wakaApi"
)

type Item struct {
	Date string
	Data interface{}
}

func main() {
	data := wakaApi.TransformData()
	dailyStats := processData.DailyTotal(data)
	langStats := processData.LanguageSummary(data)
	utils.PrettyPrint(dailyStats)
	utils.PrettyPrint(langStats.Durations)
	utils.PrettyPrint(langStats.Percentages)

}

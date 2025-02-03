package main

import (
	"log"
	"time"

	"github.com/go-echarts/snapshot-chromedp/render"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mtanzim/guac/plotData"
	"github.com/mtanzim/guac/server/services"
)

const filePath = "guac-pie-plot-test.png"

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

func jobFn(months int, topK int) error {
	dr := getDateRange(months)
	log.Printf("plotting data for top %d languages with range %s to %s", topK, dr.start, dr.end)
	rv := services.DataService(dr.start, dr.end)
	pie := plotData.LanguagePieMinimal(rv.LangStats, rv.StartDate, rv.EndDate, int64(topK))
	err := render.MakeChartSnapshot(pie.RenderContent(), filePath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := jobFn(12, 7)
	if err != nil {
		log.Fatal(err)
	}
}

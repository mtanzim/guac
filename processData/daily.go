package processData

import (
	"sort"
	"time"

	"github.com/mtanzim/guac/dynamo"
)

type DailyStat struct {
	Date string `json:"date"`
	// minutes
	Duration float64 `json:"minutes"`
}

func DailyTotal(input []dynamo.Item) []DailyStat {
	var dailyStat []DailyStat
	for _, v := range input {
		switch vv := v.Data.(type) {
		case map[string]interface{}:
			grandTotal := vv["grand_total"].(map[string]interface{})
			grandTotalInSeconds := grandTotal["total_seconds"].(float64)
			grandTotalInMins := grandTotalInSeconds / 60.0
			dailyStat = append(dailyStat, DailyStat{v.Date, grandTotalInMins})
		default:
			// Do nothing
		}
	}
	sort.Slice(dailyStat, func(i, j int) bool {
		layout := "2006-01-02"
		prevDate, _ := time.Parse(layout, dailyStat[i].Date)
		curDate, _ := time.Parse(layout, dailyStat[j].Date)
		return prevDate.Before(curDate)
	})
	return dailyStat
}

func GetDateRange(input []DailyStat) (string, string) {
	return input[0].Date, input[len(input)-1].Date
}

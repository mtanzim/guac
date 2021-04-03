package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/processData"
)

type RV struct {
	DailyStats []processData.DailyStat  `json:"dailyDuration"`
	LangStats  processData.LanguageStat `json:"languageStats"`
}

func dataService(start, end string) *RV {
	data := dynamo.GetData(start, end)
	dailyStats := processData.DailyTotal(data)
	langStats := processData.LanguageSummary(data)
	return &RV{dailyStats, langStats}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Conten-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	rv := dataService("2020-03-22", "2022-03-25")
	if err := json.NewEncoder(w).Encode(rv); err != nil {
		log.Fatalln(err)
	}
}

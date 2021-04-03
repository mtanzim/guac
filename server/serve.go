package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/processData"
)

type RV struct {
	DailyStats []processData.DailyStat  `json:"dailyDuration"`
	LangStats  processData.LanguageStat `json:"languageStats"`
	StartDate  string                   `json:"startDate"`
	EndDate    string                   `json:"endDate"`
}

func dataService(start, end string) *RV {
	data := dynamo.GetData(start, end)
	dailyStats := processData.DailyTotal(data)
	actualStart, actualEnd := processData.GetDateRange(dailyStats)
	langStats := processData.LanguageSummary(data)
	return &RV{DailyStats: dailyStats, LangStats: langStats, StartDate: actualStart, EndDate: actualEnd}
}

func validateQueryDate(start, end string) error {
	dateLayout := "2006-01-02"
	timeStart, err := time.Parse(dateLayout, start)
	if err != nil {
		return errors.New("Invalid start date")
	}

	timeEnd, err := time.Parse(dateLayout, end)
	if err != nil {
		return errors.New("Invalid end date")
	}

	if timeEnd.Before(timeStart) {
		return errors.New("End date is before start date")
	}

	return nil
}

func DataController(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqStart := req.URL.Query().Get("start")
	reqEnd := req.URL.Query().Get("end")
	if err := validateQueryDate(reqStart, reqEnd); err != nil {
		log.Println(err)
		errorRv := struct {
			Error string `json:"error"`
		}{err.Error()}
		if err := json.NewEncoder(w).Encode(errorRv); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	rv := dataService(reqStart, reqEnd)
	if err := json.NewEncoder(w).Encode(rv); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

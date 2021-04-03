package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/plotData"
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

func handlerError(w http.ResponseWriter, err error) {
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

func PlotController(w http.ResponseWriter, req *http.Request) {

	reqStart := req.URL.Query().Get("start")
	reqEnd := req.URL.Query().Get("end")
	if err := validateQueryDate(reqStart, reqEnd); err != nil {
		handlerError(w, err)
		return
	}

	reqType := req.URL.Query().Get("type")
	switch reqType {
	case "dailyBar":
		rv := dataService(reqStart, reqEnd)
		bar := plotData.DailyBarChart(rv.DailyStats, rv.StartDate, rv.EndDate)
		bar.Render(w)
	case "languagePie":
		rv := dataService(reqStart, reqEnd)
		pie := plotData.LanguagePie(rv.LangStats, rv.StartDate, rv.EndDate)
		pie.Render(w)
	case "all":
		rv := dataService(reqStart, reqEnd)
		page := plotData.Page(rv.DailyStats, rv.LangStats, rv.StartDate, rv.EndDate)
		page.Renderer.Render(w)
	default:
		handlerError(w, errors.New("Invalid chart type"))
	}

}

func DataController(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqStart := req.URL.Query().Get("start")
	reqEnd := req.URL.Query().Get("end")
	if err := validateQueryDate(reqStart, reqEnd); err != nil {
		handlerError(w, err)
		return
	}
	rv := dataService(reqStart, reqEnd)
	if err := json.NewEncoder(w).Encode(rv); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

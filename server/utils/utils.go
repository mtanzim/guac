package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mtanzim/guac/processData"
)

type RV struct {
	DailyStats []processData.DailyStat  `json:"dailyDuration"`
	LangStats  processData.LanguageStat `json:"languageStats"`
	ProjStats  processData.ProjectStat  `json:"projectStats"`
	StartDate  string                   `json:"startDate"`
	EndDate    string                   `json:"endDate"`
}

func ValidateQueryDate(start, end string) error {
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

func HandlerError(w http.ResponseWriter, err error) {
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

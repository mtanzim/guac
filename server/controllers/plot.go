package controllers

import (
	"errors"
	"net/http"
	"os"

	"github.com/mtanzim/guac/plotData"
	"github.com/mtanzim/guac/server/services"
	"github.com/mtanzim/guac/server/utils"
)

func PlotController(w http.ResponseWriter, req *http.Request) {

	reqStart := req.URL.Query().Get("start")
	reqEnd := req.URL.Query().Get("end")

	if reqStart == "" || reqEnd == "" {
		reqStart, reqEnd = os.Getenv("START"), os.Getenv("END")

		if reqStart == "" || reqEnd == "" {
			utils.HandlerError(w, errors.New("Please configure default start and end dates"))
			return
		}

	} else {
		if err := utils.ValidateQueryDate(reqStart, reqEnd); err != nil {
			utils.HandlerError(w, err)
			return
		}
	}

	reqType := req.URL.Query().Get("type")
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	switch reqType {
	case "dailyBar":
		rv := services.DataService(reqStart, reqEnd)
		bar := plotData.DailyBarChart(rv.DailyStats, rv.StartDate, rv.EndDate)
		bar.Render(w)
	case "projectBar":
		rv := services.DataService(reqStart, reqEnd)
		bar := plotData.ProjectBarChart(rv.ProjStats, rv.StartDate, rv.EndDate)
		bar.Render(w)
	case "languagePie":
		rv := services.DataService(reqStart, reqEnd)
		pie := plotData.LanguagePie(rv.LangStats, rv.StartDate, rv.EndDate)

		pie.Render(w)
	default:
		rv := services.DataService(reqStart, reqEnd)
		page := plotData.Page(rv.DailyStats, rv.LangStats, rv.ProjStats, rv.StartDate, rv.EndDate)
		page.Renderer.Render(w)
	}

}

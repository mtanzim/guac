package controllers

import (
	"errors"
	"net/http"
	"os"

	"github.com/mtanzim/guac/plotData"
	"github.com/mtanzim/guac/server/services"
	"github.com/mtanzim/guac/server/utils"
)

func RootController(w http.ResponseWriter, req *http.Request) {

	start, end := os.Getenv("START"), os.Getenv("END")

	if start == "" || end == "" {
		utils.HandlerError(w, errors.New("Please configure default start and end dates"))
	}

	rv := services.DataService(start, end)
	page := plotData.Page(rv.DailyStats, rv.LangStats, rv.ProjStats, rv.StartDate, rv.EndDate)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	page.Renderer.Render(w)

}

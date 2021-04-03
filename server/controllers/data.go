package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mtanzim/guac/server/services"
	"github.com/mtanzim/guac/server/utils"
)

func DataController(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqStart := req.URL.Query().Get("start")
	reqEnd := req.URL.Query().Get("end")
	if err := utils.ValidateQueryDate(reqStart, reqEnd); err != nil {
		utils.HandlerError(w, err)
		return
	}
	rv := services.DataService(reqStart, reqEnd)
	if err := json.NewEncoder(w).Encode(rv); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

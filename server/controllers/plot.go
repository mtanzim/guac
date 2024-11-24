package controllers

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/go-echarts/snapshot-chromedp/render"
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

func PlotImageController(w http.ResponseWriter, req *http.Request) {

	reqStart := req.URL.Query().Get("start")
	reqEnd := req.URL.Query().Get("end")

	topK := int64(5)
	topKs := req.URL.Query().Get("topK")
	if topKs != "" {
		k, err := strconv.ParseInt(topKs, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("error: %v\n", err)
			w.Write([]byte("top k must be an int"))
			return
		}
		topK = k
	}

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

	w.Header().Set("Content-Type", "application/octet-stream")
	rv := services.DataService(reqStart, reqEnd)
	pie := plotData.LanguagePieMinimal(rv.LangStats, rv.StartDate, rv.EndDate, topK)

	bs := pie.RenderContent()
	imgBs, err := makeSnapshot(bs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(imgBs)

}

// copy pasta from internal lib
func makeSnapshot(bs []byte) ([]byte, error) {
	path, file := filepath.Split("temp.png")
	suffix := filepath.Ext(file)[1:]
	fileName := file[0 : len(file)-len(suffix)-1]
	config := &render.SnapshotConfig{
		RenderContent: bs,
		Path:          path,
		FileName:      fileName,
		Suffix:        suffix,
		Quality:       1,
		KeepHtml:      false,
		Timeout:       0,
	}
	content := config.RenderContent
	quality := config.Quality
	keepHtml := config.KeepHtml
	htmlPath := config.HtmlPath
	timeout := config.Timeout

	if htmlPath == "" {
		htmlPath = path
	}

	if !filepath.IsAbs(path) {
		path, _ = filepath.Abs(path)
	}

	if !filepath.IsAbs(htmlPath) {
		htmlPath, _ = filepath.Abs(htmlPath)
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if timeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	htmlFullPath := filepath.Join(htmlPath, fileName+"."+render.HTML)

	if !keepHtml {
		defer func() {
			err := os.Remove(htmlFullPath)
			if err != nil {
				log.Printf("Failed to delete the file(%s), err: %s\n", htmlFullPath, err)
			}
		}()
	}

	err := os.WriteFile(htmlFullPath, content, 0o644)
	if err != nil {
		return []byte{}, err
	}

	if quality < 1 {
		quality = 1
	}

	var imgContent []byte
	executeJS := fmt.Sprintf(render.CanvasJs, suffix, quality)
	pagePath := fmt.Sprintf("%s%s", render.FileProtocol, htmlFullPath)

	if true != config.MultiCharts {
		imgContent, err = snapshotSingleChart(ctx, pagePath, executeJS)
	} else {
		imgContent, err = snapshotMultiCharts(ctx, pagePath, quality)
	}

	if err != nil {
		return []byte{}, err
	}

	return imgContent, nil

}

func snapshotSingleChart(ctx context.Context, pagePath string, executeJS string) ([]byte, error) {
	var base64Data string
	var imageContent []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(pagePath),
		chromedp.WaitVisible(render.EchartsInstanceDom, chromedp.ByQuery),
		chromedp.Evaluate(executeJS, &base64Data),
	)

	if err != nil {
		return nil, err
	}
	imageContent, err = base64.StdEncoding.DecodeString(strings.Split(base64Data, ",")[1])
	return imageContent, err

}

func snapshotMultiCharts(ctx context.Context, pagePath string, quality int) ([]byte, error) {
	var imageContent []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(pagePath),
		chromedp.WaitVisible(render.EchartsInstanceDom, chromedp.ByQuery),
		chromedp.FullScreenshot(&imageContent, quality),
	)

	return imageContent, err

}

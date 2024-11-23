package plotData

import (
	"math"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mtanzim/guac/processData"
)

func DailyBarChart(dailyStats []processData.DailyStat, start, end string) *charts.Bar {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title:    "My coding activity",
			Subtitle: start + " to " + end,
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Date"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Duration (Hours)"}))

	var xs []string
	var ys []opts.BarData
	for _, v := range dailyStats {
		xs = append(xs, v.Date)
		h := v.Duration / 60.0
		h = math.Round(h*100) / 100
		ys = append(ys, opts.BarData{Value: h})
	}
	seriesOpts := charts.WithEmphasisOpts(opts.Emphasis{Label: &opts.Label{Show: opts.Bool(true), Color: "black", Position: "top", Formatter: "{c} hours on {b}"}})

	// Put data into instance
	bar.SetXAxis(xs).
		AddSeries("Duration (Hours)", ys, seriesOpts)
	// // Where the magic happens
	// f, _ := os.Create("bar.html")
	// bar.Render(f)
	return bar
}

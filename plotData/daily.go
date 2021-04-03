package plotData

import (
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
		charts.WithYAxisOpts(opts.YAxis{Name: "Duration (minutes)"}))

	var xs []string
	var ys []opts.BarData
	for _, v := range dailyStats {
		xs = append(xs, v.Date)
		ys = append(ys, opts.BarData{Value: v.Duration})
	}

	// Put data into instance
	bar.SetXAxis(xs).
		AddSeries("Duration (minutes)", ys)
	// // Where the magic happens
	// f, _ := os.Create("bar.html")
	// bar.Render(f)
	return bar
}

package plotData

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mtanzim/guac/processData"
)

func ProjectBarChart(projStats processData.ProjectStat, start, end string) *charts.Bar {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title:    "My Projects",
			Subtitle: start + " to " + end,
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Project"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Duration (hours)"}))

	var xs []string
	var ys []opts.BarData
	for _, v := range projStats.Durations {
		xs = append(xs, v.Name)
		ys = append(ys, opts.BarData{Value: v.Duration / 60.0})
	}

	bar.SetXAxis(xs).
		AddSeries("Duration (hours)", ys)
	return bar
}

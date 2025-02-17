package plotData

import (
	"math"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mtanzim/guac/processData"
)

func LanguagePieMinimal(langStats processData.LanguageStat, start, end string, topK int64) *charts.Pie {
	colors := NewColors()
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title: "",
			Left:  "center",
		},
	),
		charts.WithLegendOpts(opts.Legend{Orient: "vertical", Show: opts.Bool(false), Left: "left"}),
		charts.WithAnimation(false),
	)

	topKPcts := processData.KLanguagePct(processData.SynonymizeLanguagePcts(langStats.Percentages), int(topK))
	var items []opts.PieData
	for _, v := range topKPcts {
		items = append(items,
			opts.PieData{
				Name:      v.Name,
				Value:     v.Pct,
				ItemStyle: &opts.ItemStyle{Color: colors.GetColor(v.Name)},
			})
	}

	pie.AddSeries("pie", items).SetSeriesOptions(
		charts.WithPieChartOpts(opts.PieChart{
			Radius: []string{"40%", "75%"},
		}),
		charts.WithLabelOpts(opts.Label{
			Color:      "#85d4e6",
			FontWeight: "200",
			FontSize:   12,
			FontFamily: "monospace",
		}),
	)
	return pie
}

func LanguagePie(langStats processData.LanguageStat, start, end string) *charts.Pie {
	colors := NewColors()
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title:    "Languages used",
			Subtitle: start + " to " + end,
			Left:     "center",
		},
	),
		charts.WithLegendOpts(opts.Legend{Orient: "vertical", Show: opts.Bool(true), Left: "left"}),
	)

	var items []opts.PieData
	for _, v := range langStats.Percentages {
		items = append(items,
			opts.PieData{
				Name:      v.Name,
				Value:     v.Pct,
				ItemStyle: &opts.ItemStyle{Color: colors.GetColor(v.Name)},
			})
	}

	pie.AddSeries("pie", items).SetSeriesOptions(
		charts.WithPieChartOpts(opts.PieChart{
			Radius: []string{"40%", "75%"},
		}),
	)
	return pie
}

func LanguageBarChart(langStats processData.LanguageStat, start, end string) *charts.Bar {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title:    "Language Durations",
			Subtitle: start + " to " + end,
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Language", AxisLabel: &opts.AxisLabel{Rotate: 60}}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Duration (hours)"}))

	var xs []string
	var ys []opts.BarData
	for _, v := range langStats.Durations {
		xs = append(xs, v.Name)
		h := v.Duration / 60.0
		h = math.Round(h*100) / 100
		ys = append(ys, opts.BarData{Value: h})
	}

	seriesOpts := charts.WithEmphasisOpts(opts.Emphasis{Label: &opts.Label{Show: opts.Bool(true), Color: "black", Position: "top", Formatter: "{c} hours on {b}"}})

	bar.SetXAxis(xs).
		AddSeries("Duration (hours)", ys, charts.SeriesOpts(seriesOpts))
	return bar
}

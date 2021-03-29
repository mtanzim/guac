package plotData

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mtanzim/guac/processData"
)

func LanguagePie(langPcts []processData.LangPct, start, end string) {
	colors := NewColors()
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title:    "Languages used",
			Subtitle: start + " to " + end,
			Left:     "center",
		},
	),
		charts.WithLegendOpts(opts.Legend{Orient: "vertical", Show: true, Left: "left"}),
	)

	var items []opts.PieData
	for _, v := range langPcts {
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
	f, _ := os.Create("pie.html")
	pie.Render(f)
}

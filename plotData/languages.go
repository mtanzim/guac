package plotData

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mtanzim/guac/processData"
)

func generatePieItems() []opts.PieData {

	var (
		itemCntPie = 4
		seasons    = []string{"Spring", "Summer", "Autumn ", "Winter"}
	)
	items := make([]opts.PieData, 0)
	for i := 0; i < itemCntPie; i++ {
		items = append(items, opts.PieData{Name: seasons[i], Value: rand.Intn(100)})
	}
	return items
}

func LanguagePie(langPcts []processData.LangPct) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic pie example"}),
	)
	pie.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title: "Languages used",
			Left:  "center",
		}))

	var items []opts.PieData
	for _, v := range langPcts {
		items = append(items, opts.PieData{Name: v.Name, Value: v.Pct})
	}

	pie.AddSeries("pie", items)
	f, _ := os.Create("pie.html")
	pie.Render(f)
}

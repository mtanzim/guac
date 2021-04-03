package plotData

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/mtanzim/guac/processData"
)

func Page(dailyStats []processData.DailyStat, langStats processData.LanguageStat, start, end string) *components.Page {
	page := components.NewPage()
	page.SetLayout(components.PageFlexLayout)
	line := DailyBarChart(dailyStats, start, end)
	pie := LanguagePie(langStats, start, end)
	page.AddCharts(pie, line)
	return page
}

package plotData

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/mtanzim/guac/processData"
)

func Page(dailyStats []processData.DailyStat, langStats processData.LanguageStat, projStats processData.ProjectStat, start, end string) *components.Page {
	page := components.NewPage()
	page.SetLayout(components.PageFlexLayout)
	barDaily := DailyBarChart(dailyStats, start, end)
	barProjects := ProjectBarChart(projStats, start, end)
	pie := LanguagePie(langStats, start, end)
	barLang := LanguageBarChart(langStats, start, end)
	page.AddCharts(pie, barLang, barProjects, barDaily)
	return page
}

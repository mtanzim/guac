package services

import (
	"github.com/mtanzim/guac/dynamo"
	"github.com/mtanzim/guac/processData"
	"github.com/mtanzim/guac/server/utils"
)

func DataService(start, end string) *utils.RV {
	data := dynamo.GetData(start, end)
	dailyStats := processData.DailyTotal(data)
	actualStart, actualEnd := processData.GetDateRange(dailyStats)
	langStats := processData.LanguageSummary(data)
	projectStats := processData.ProjectSummary(data)
	return &utils.RV{DailyStats: dailyStats, LangStats: langStats, ProjStats: projectStats, StartDate: actualStart, EndDate: actualEnd}
}

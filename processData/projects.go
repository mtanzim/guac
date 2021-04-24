package processData

import (
	"sort"

	"github.com/mtanzim/guac/dynamo"
)

type ProjectStat struct {
	Durations []ProjectDur `json:"durations"`
}

type ProjectDur struct {
	Name string `json:"project"`
	// minutes
	Duration float64 `json:"minutes"`
}

func projectDuration(input []dynamo.Item) map[string]float64 {
	projectSummary := make(map[string]float64)
	for _, v := range input {
		switch vv := v.Data.(type) {
		case map[string]interface{}:

			if vv["projects"] == nil {
				continue
			}
			projects := vv["projects"].([]interface{})

			for _, lang := range projects {
				switch ll := lang.(type) {
				case map[string]interface{}:
					projectName := ll["name"].(string)
					projectDurationInSec := ll["total_seconds"].(float64)
					projectDurationInMin := projectDurationInSec / 60.0
					if val, ok := projectSummary[projectName]; ok {
						projectSummary[projectName] = val + projectDurationInMin
					} else {
						projectSummary[projectName] = projectDurationInMin
					}
				}

			}
		}
	}
	return projectSummary
}

func transformProjectDurationsMap(durations map[string]float64) []ProjectDur {
	var durationsSlc []ProjectDur
	for k, v := range durations {
		durationsSlc = append(durationsSlc, ProjectDur{k, v})
	}
	sort.Slice(durationsSlc, func(i, j int) bool {
		return durationsSlc[i].Duration > durationsSlc[j].Duration
	})
	return durationsSlc
}

func ProjectSummary(input []dynamo.Item) ProjectStat {
	durations := projectDuration(input)
	durationsSlc := transformProjectDurationsMap(durations)
	return ProjectStat{durationsSlc}
}

package processData

import (
	"errors"
	"log"
	"math"
	"sort"
)

type LanguageStat struct {
	Durations   []LangDur
	Percentages []LangPct
}

type LangDur struct {
	Name string
	// minutes
	Duration float64
}

type LangPct struct {
	Name string
	Pct  float64
}

// TODO: this shit ugly, is it idiomatic?
func languageDuration(input map[string]interface{}) map[string]float64 {
	languageSummary := make(map[string]float64)
	for _, v := range input {
		switch vv := v.(type) {
		case map[string]interface{}:
			languages := vv["languages"].([]interface{})
			for _, lang := range languages {
				switch ll := lang.(type) {
				case map[string]interface{}:
					languageName := ll["name"].(string)
					languageDurationInSec := ll["total_seconds"].(float64)
					languageDurationInMin := languageDurationInSec / 60.0
					if val, ok := languageSummary[languageName]; ok {
						languageSummary[languageName] = val + languageDurationInMin
					} else {
						languageSummary[languageName] = languageDurationInMin
					}
				}

			}
		}
	}
	return languageSummary
}

func languagePct(durations map[string]float64) ([]LangPct, error) {
	var percentages []LangPct
	totalDur := 0.0
	for _, v := range durations {
		totalDur += v
	}
	pctTotal := 0.0
	for k, v := range durations {
		curPct := (v / totalDur) * 100.0
		pctTotal += curPct
		percentages = append(percentages, LangPct{k, curPct})

	}
	if epsilon := 0.000000001; math.Abs(pctTotal-100) > epsilon {
		return nil, errors.New("Pct calculation errors")
	}
	sort.Slice(percentages, func(i, j int) bool {
		return percentages[i].Pct > percentages[j].Pct
	})
	return percentages, nil
}

func transformDurationsMap(durations map[string]float64) []LangDur {
	var durationsSlc []LangDur
	for k, v := range durations {
		durationsSlc = append(durationsSlc, LangDur{k, v})
	}
	sort.Slice(durationsSlc, func(i, j int) bool {
		return durationsSlc[i].Duration > durationsSlc[j].Duration
	})
	return durationsSlc
}

func LanguageSummary(input map[string]interface{}) LanguageStat {
	durations := languageDuration(input)
	pct, err := languagePct(durations)
	if err != nil {
		log.Panicln(err.Error())
	}
	durationsSlc := transformDurationsMap(durations)
	return LanguageStat{durationsSlc, pct}
}

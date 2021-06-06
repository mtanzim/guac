package processData

import (
	"log"
	"sort"

	"github.com/mtanzim/guac/firestoreClient"
)

type LanguageStat struct {
	Durations   []LangDur `json:"durations"`
	Percentages []LangPct `json:"percentages"`
}

type LangDur struct {
	Name string `json:"language"`
	// minutes
	Duration float64 `json:"minutes"`
}

type LangPct struct {
	Name string  `json:"language"`
	Pct  float64 `json:"percentage"`
}

const MAX_LANG_COUNT = 5

// TODO: this shit ugly, is it idiomatic?
func languageDuration(input []firestoreClient.Item) map[string]float64 {
	languageSummary := make(map[string]float64)
	for _, v := range input {
		switch vv := v.Data.(type) {
		case map[string]interface{}:

			if vv["languages"] == nil {
				continue
			}
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

func CleanLangPct(sortedLangPct []LangPct) []LangPct {
	i := 0
	totalPct := 0.0
	undesiredLangsSet := make(map[string]bool)
	for _, v := range [...]string{"JSON", "Other", "HTML", "YAML", "Markdown"} {
		undesiredLangsSet[v] = true
	}
	var topLangPct []LangPct
	for _, pct := range sortedLangPct {
		if i == MAX_LANG_COUNT {
			break
		}
		if _, ok := undesiredLangsSet[pct.Name]; ok {
			continue
		}
		topLangPct = append(topLangPct, pct)
		totalPct += pct.Pct
		i += 1
	}
	remainingPct := 100 - totalPct
	if i < len(sortedLangPct)-1 {
		topLangPct = append(topLangPct, LangPct{Name: "Other", Pct: remainingPct})
	}
	return topLangPct

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
	sort.Slice(percentages, func(i, j int) bool {
		return percentages[i].Pct > percentages[j].Pct
	})

	return percentages, nil
}

func transformLangDurationsMap(durations map[string]float64) []LangDur {
	var durationsSlc []LangDur
	for k, v := range durations {
		durationsSlc = append(durationsSlc, LangDur{k, v})
	}
	sort.Slice(durationsSlc, func(i, j int) bool {
		return durationsSlc[i].Duration > durationsSlc[j].Duration
	})
	return durationsSlc
}

func LanguageSummary(input []firestoreClient.Item) LanguageStat {
	durations := languageDuration(input)
	pct, err := languagePct(durations)
	if err != nil {
		log.Fatalln(err.Error())
	}
	durationsSlc := transformLangDurationsMap(durations)
	return LanguageStat{durationsSlc, pct}
}

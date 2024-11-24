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
	var languages []interface{}
	for _, v := range input {
		switch vv := v.Data.(type) {
		case map[string]interface{}:
			if vv["languages"] != nil {
				dailyLanguages := vv["languages"].([]interface{})
				languages = append(languages, dailyLanguages...)
				break
			}
		}
	}

	languageSummary := make(map[string]float64)
	if languages == nil {
		return languageSummary
	}
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

const (
	RestLang   = "Rest"
	MarkupLang = "Markup and config"
)

func SynonimizeLanguagePcts(pcts []LangPct) []LangPct {

	synonyms := map[string]string{
		"JSX":             "JavaScript",
		"JSON":            "JavaScript",
		"TSX":             "Typescript",
		"Astro":           "Typescript",
		"YAML":            MarkupLang,
		"HTML":            MarkupLang,
		"Text":            MarkupLang,
		"XML":             MarkupLang,
		"INI":             MarkupLang,
		"Smarty":          MarkupLang,
		"Markdown":        MarkupLang,
		"Docker":          MarkupLang,
		"TOML":            MarkupLang,
		"Makefile":        MarkupLang,
		"Image (svg)":     MarkupLang,
		"Git Config":      MarkupLang,
		"TSConfig":        MarkupLang,
		"Protocol Buffer": MarkupLang,
		"Terraform":       MarkupLang,
		"CSV":             MarkupLang,
		"Crontab":         MarkupLang,
		"Git":             MarkupLang,
		"Other":           MarkupLang,
	}
	pctHm := map[string]float64{}
	for _, v := range pcts {
		pctHm[v.Name] = v.Pct
	}
	pctHmPrime := map[string]float64{}
	for k, v := range pctHm {
		synK, ok := synonyms[k]
		if !ok {
			pctHmPrime[k] = v
			continue
		}
		curPctVal, alreadyThere := pctHmPrime[synK]
		if alreadyThere {
			pctHmPrime[synK] = v + curPctVal
			continue
		}
		pctHmPrime[synK] = v

	}
	res := []LangPct{}
	for k, v := range pctHmPrime {
		res = append(res, LangPct{Name: k, Pct: v})
	}
	return res

}

func KLanguagePct(pcts []LangPct, topK int) []LangPct {
	sort.Slice(pcts, func(i, j int) bool {
		return pcts[i].Pct > pcts[j].Pct
	})
	restLangPct := &LangPct{
		Name: RestLang,
		Pct:  0.0,
	}
	res := []LangPct{}
	for i, v := range pcts {
		if i < topK {
			res = append(res, v)
			continue
		}
		restLangPct.Pct += v.Pct
	}
	if restLangPct.Pct > 0 {
		res = append(res, *restLangPct)
	}
	return res

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
		// TODO: why panic here?
		log.Fatalln(err.Error())
	}
	durationsSlc := transformLangDurationsMap(durations)
	return LanguageStat{durationsSlc, pct}
}

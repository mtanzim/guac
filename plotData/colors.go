package plotData

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type LanguageColors struct {
	Colors map[string]string
}

const DEFAULT_COLOR = "#800080"

func (c LanguageColors) GetColor(name string) string {
	if curColor, ok := c.Colors[name]; ok {
		log.Println(curColor)
		return curColor
	} else {
		return DEFAULT_COLOR
	}
}

func NewColors() *LanguageColors {

	// Open our jsonFile
	jsonFile, err := os.Open("colors.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panicln(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	colorsJson, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(colorsJson), &result)

	colors := make(map[string]string)
	for k, v := range result {
		colors[k] = v.(string)
	}
	return &LanguageColors{colors}

}

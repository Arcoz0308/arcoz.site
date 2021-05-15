package exercises

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func loadLanguages() {
	content, err := ioutil.ReadFile("asset/json/simpleCodeLanguagesList.json")
	if err != nil {
		log.Fatal(err)
	}
	langs := make(map[string]map[string][]string)
	err = json.Unmarshal(content, &langs)
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range langs {
		languages[key] = key
		for _, v := range val["aliases"] {
			languages[v] = key
		}
	}
}

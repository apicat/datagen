package datagen

import (
	"embed"
	"encoding/json"
	"strings"
)

//go:embed locales/*.json
var i18nDataDir embed.FS
var localesDatas map[string]*locale

const defaultLocal = "en"

type locale struct {
	Text    *i18nText    `json:"text,omitempty"`
	People  *i18nPeople  `json:"people,omitempty"`
	Address *i18nAddress `json:"address,omitempty"`
}

func init() {
	ls, err := i18nDataDir.ReadDir("locales")
	if err != nil {
		panic(err)
	}
	localesDatas = make(map[string]*locale)
	for _, v := range ls {
		if v.IsDir() {
			continue
		}
		raw, err := i18nDataDir.ReadFile("locales/" + v.Name())
		if err != nil {
			panic(err)
		}
		var data locale
		if err := json.Unmarshal(raw, &data); err != nil {
			panic(err)
		}
		data.Text._charRuns = []rune(data.Text.Chars)
		localesDatas[strings.Split(v.Name(), ".")[0]] = &data
	}
}

func GetLocale(lang ...string) *locale {
	if len(lang) > 0 {
		if v, ok := localesDatas[pick(lang)]; ok {
			return v
		}
	}
	return localesDatas[defaultLocal]
}

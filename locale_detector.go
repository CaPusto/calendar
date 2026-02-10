// locale_detector.go
package main

import (
	"strings"

	golocale "github.com/jeandeaual/go-locale"
)

// langToKey — сопоставление основного кода языка (ISO 639-1) → ключ локали в проекте
var langToKey = map[string]string{
	"hy": "Armenia (HY)",       // армянский
	"en": "USA (EN)",           // английский — по умолчанию USA
	"az": "Azerbaijan (AZ)",
	"be": "Belarus (BE)",
	"pt": "Brazil (PT)",        // португальский (Бразилия)
	"zh": "China (ZH)",         // китайский
	"et": "Estonia (ET)",
	"fi": "Finland (FI)",
	"fr": "France (FR)",        // французский — по умолчанию Франция
	"ka": "Georgia (KA)",
	"de": "Germany (DE)",
	"hi": "India (HI)",         // хинди
	"fa": "Iran (FA)",          // персидский
	"he": "Israel (HE)",        // иврит
	"it": "Italy (IT)",
	"ja": "Japan (JA)",
	"kk": "Kazakhstan (KK)",
	"ky": "Kyrgyzstan (KY)",
	"lv": "Latvia (LV)",
	"lt": "Lithuania (LT)",
	"ms": "Malaysia (MS)",
	"es": "Spain (ES)",         // испанский — по умолчанию Испания
	"ro": "Moldova (RO)",
	"pl": "Poland (PL)",
	"ru": "Russia (RU)",
	"tg": "Tajikistan (TG)",
	"tr": "Turkey (TR)",
	"tk": "Turkmenistan (TK)",
	"uk": "Ukraine (UK)",
	"uz": "Uzbekistan (UZ)",
	"vi": "Vietnam (VI)",
}

// detectPreferredCountryKey возвращает наиболее подходящий ключ локали
// на основе языка системы пользователя
func detectPreferredCountryKey() string {
	locales, err := golocale.GetLocales()
	if err != nil || len(locales) == 0 {
		return fallbackKey()
	}

	// Самый приоритетный язык пользователя
	preferred := strings.ToLower(locales[0]) // например: "ru-ru", "de-de", "en-us", "he-il"

	// Отделяем код языка от региона
	langPart := strings.Split(preferred, "-")[0]

	// Прямое сопоставление
	if key, ok := langToKey[langPart]; ok {
		// Специальные уточнения по региону (если нужно)
		if langPart == "en" {
			region := ""
			if len(strings.Split(preferred, "-")) > 1 {
				region = strings.ToLower(strings.Split(preferred, "-")[1])
			}
			if region == "gb" || region == "uk" {
				if containsKey("United Kingdom (EN)") {
					return "United Kingdom (EN)"
				}
			}
			if region == "au" && containsKey("Australia (EN)") {
				return "Australia (EN)"
			}
			if region == "nz" && containsKey("New Zealand (EN)") {
				return "New Zealand (EN)"
			}
			if region == "ca" && containsKey("Canada (EN)") {
				return "Canada (EN)"
			}
			// по умолчанию USA
			return "USA (EN)"
		}

		if langPart == "fr" {
			region := ""
			if len(strings.Split(preferred, "-")) > 1 {
				region = strings.ToLower(strings.Split(preferred, "-")[1])
			}
			if region == "ca" && containsKey("Canada (FR)") {
				return "Canada (FR)"
			}
			return "France (FR)"
		}

		return key
	}

	// Если язык не найден в мапе — fallback
	return fallbackKey()
}

// fallbackKey — запасной вариант
func fallbackKey() string {
	// Можно выбрать "USA (EN)" как наиболее нейтральный международный вариант
	// или "Russia (RU)", если проект ориентирован на русскоязычную аудиторию
	return "USA (EN)"
}

// containsKey — проверяет, существует ли ключ в глобальном списке countryKeys
func containsKey(key string) bool {
	for _, k := range countryKeys {
		if k == key {
			return true
		}
	}
	return false
}
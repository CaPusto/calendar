package main

import (
	"encoding/json"
	"os"
	"strings"
)

type Locale struct {
	Title, YearLabel, WeekDayLabel, OutputLabel, CountryLabel, GoLabel, HelpMsg, SuccessMsg, ErrorMsg string
	Months                                                                                            []string
	HeaderSun, HeaderMon                                                                              string
	SunVal, MonVal, ConsoleVal, FileVal                                                               string
	DefaultMon                                                                                        bool
}

var countryKeys = []string{
	"Albania (SQ)", "Andorra (CA)", "Armenia (HY)", "Australia (EN)", "Austria (DE)", "Azerbaijan (AZ)",
	"Belarus (BE)", "Belgium (NL)", "Bosnia and Herzegovina (BS)", "Brazil (PT)", "Bulgaria (BG)",
	"Canada (EN)", "Canada (FR)", "China (ZH)", "Cyprus (EL)", "Czech Republic (CS)", "Denmark (DA)",
	"Estonia (ET)", "Finland (FI)", "France (FR)", "Georgia (KA)", "Germany (DE)", "Greece (EL)",
	"Hungary (HU)", "Iceland (IS)", "India (HI)", "Iran (FA)", "Ireland (EN)", "Israel (HE)", "Italy (IT)",
	"Japan (JA)", "Kazakhstan (KK)", "Kyrgyzstan (KY)", "Latvia (LV)", "Liechtenstein (DE)",
	"Lithuania (LT)", "Luxembourg (LB)", "Malaysia (MS)", "Malta (MT)", "Moldova (RO)", "Monaco (FR)",
	"Montenegro (ME)", "Netherlands (NL)", "New Zealand (EN)", "North Macedonia (MK)", "Norway (NO)",
	"Poland (PL)", "Portugal (PT)", "Romania (RO)", "Russia (RU)", "San Marino (IT)", "Serbia (SR)",
	"Slovakia (SK)", "Slovenia (SL)", "Spain (ES)", "Sweden (SV)", "Switzerland (DE)", "Tajikistan (TG)",
	"Turkey (TR)", "Ukraine (UK)", "United Kingdom (EN)", "USA (EN)", "Uzbekistan (UZ)", "Vatican (IT)",
	"Vietnam (VI)",
}

var locales = map[string]Locale{
	"Russia (RU)": {
		Title: "=== НАСТРОЙКИ КАЛЕНДАРЯ ===", YearLabel: "Календарь на:", WeekDayLabel: "Первый день недели:",
		OutputLabel: "Вывод в:", CountryLabel: "Страна, Язык:", GoLabel: "Запуск:", HelpMsg: "[←/→] Изменить | [Enter] Пуск | [q] Выход",
		SuccessMsg: "\nУСПЕХ! Сохранено в файл calendar.txt\n",
		ErrorMsg:   "\nОШИБКА! Не удалось сохранить в файл calendar.txt: %v\n",
		Months:     []string{"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"},
		HeaderSun:  "Вс Пн Вт Ср Чт Пт Сб", HeaderMon: "Пн Вт Ср Чт Пт Сб Вс",
		SunVal: "Воскресенье", MonVal: "Понедельник", ConsoleVal: "Консоль", FileVal: "Файл", DefaultMon: true,
	},
	"USA (EN)": {
		Title: "=== CALENDAR SETTINGS ===", YearLabel: "Calendar for:", WeekDayLabel: "First day of the week:",
		OutputLabel: "Output to:", CountryLabel: "Country, Lang:", GoLabel: "Go:", HelpMsg: "[←/→] Change | [Enter] Go | [q] Quit",
		SuccessMsg: "\nSUCCESS! Saved to calendar.txt\n",
		ErrorMsg:   "\nERROR! Failed to save to calendar.txt: %v\n",
		Months:     []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		HeaderSun:  "Su Mo Tu We Th Fr Sa", HeaderMon: "Mo Tu We Th Fr Sa Su",
		SunVal: "Sunday", MonVal: "Monday", ConsoleVal: "Console", FileVal: "File", DefaultMon: false,
	},
}

func getLocale(key string) Locale {
	if l, ok := locales[key]; ok {
		return l
	}
	return locales["USA (EN)"] // fallback
}

func loadExtraLocales() {
	entries, err := os.ReadDir("locales")
	if err != nil {
		return 
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile("locales/" + e.Name())
		if err != nil {
			continue
		}

		var loc Locale
		if err := json.Unmarshal(data, &loc); err != nil {
			continue
		}

		key := strings.TrimSuffix(e.Name(), ".json")
		key = strings.ReplaceAll(key, "_", " ")
		locales[key] = loc
	}
}

func init() {
	loadExtraLocales()
}


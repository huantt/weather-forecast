package collector

import (
	"bytes"
	_ "embed"
	"html/template"
	"time"
	"weather_forecast/model"
)

func formatDate(date time.Time, timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	date = date.In(loc)
	return date.Format("02/01/2006")
}

func formatHour(date time.Time, timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	date = date.In(loc)
	return date.Format("15:04")
}

func todayHourlyWeatherTable(hourlyWeathers []model.Weather) string {
	var result bytes.Buffer
	tmpl, err := template.
		New("test").
		Funcs(template.FuncMap{
			"formatDate": formatDate,
			"formatHour": formatHour,
		}).
		Parse(hourlyWeatherTemplateData)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&result, map[string]any{
		"Weathers": hourlyWeathers,
	})
	if err != nil {
		panic(err)
	}
	return result.String()
}

func dailyWeatherTable(dailyWeathers []model.Weather) string {
	var result bytes.Buffer
	tmpl, err := template.
		New("test").
		Funcs(template.FuncMap{
			"formatDate": formatDate,
			"formatHour": formatHour,
		}).
		Parse(dailyWeatherTemplateData)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&result, map[string]any{
		"Weathers": dailyWeathers,
	})
	if err != nil {
		panic(err)
	}
	return result.String()
}

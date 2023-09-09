package collector

import _ "embed"

//go:embed template/hourly-forecast.md.template
var hourlyWeatherTemplateData string

//go:embed template/daily-forecast.md.template
var dailyWeatherTemplateData string

var templates = []string{
	hourlyWeatherTemplateData, dailyWeatherTemplateData,
}

package collector

import (
	"bytes"
	"context"
	_ "embed"
	"html/template"
	"os"
	"time"
	"weather_forecast/model"
	"weather_forecast/pkg/errs"
)

type Collector struct {
	weatherService WeatherService
}

func NewCollector(weatherService WeatherService) *Collector {
	return &Collector{weatherService}
}

func (c *Collector) Collect(ctx context.Context, city string, days int) error {
	weathers, err := c.weatherService.Forecast(ctx, city, days)
	if err != nil {
		return errs.Joinf(err, "[weatherService.Forecast]")
	}
	readme, err := generateReadme(weathers)
	if err != nil {
		return errs.Joinf(err, "[generateReadme]")
	}

	return os.WriteFile("README.md", []byte(*readme), 0644)
}

//go:embed data/README.md.template
var readmeTemplate string

func generateReadme(weathers []model.Weather) (*string, error) {
	tmpl, err := template.
		New("test").
		Funcs(template.FuncMap{
			"formatDate": func(date time.Time, timezone string) string {
				loc, _ := time.LoadLocation(timezone)
				date = date.In(loc)
				return date.Format("02/01/2006")
			},
			"formatHour": func(date time.Time, timezone string) string {
				loc, _ := time.LoadLocation(timezone)
				date = date.In(loc)
				return date.Format("15:04")
			},
		}).
		Parse(readmeTemplate)
	if err != nil {
		panic(err)
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, map[string]any{
		"Weathers": weathers,
	})
	if err != nil {
		return nil, err
	}
	stringResult := result.String()
	return &stringResult, nil
}

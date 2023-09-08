package collector

import (
	"context"
	"weather_forecast/model"
)

type WeatherService interface {
	Forecast(ctx context.Context, city string, days int) ([]model.Weather, error)
}

package weatherapi

import (
	"context"
	"github.com/huantt/weather-forecast/model"
	"github.com/huantt/weather-forecast/pkg/utils"
	"github.com/huantt/weather-forecast/pkg/weatherapi_com"
	"strings"
	"time"
)

type WeatherService struct {
	service *weatherapi_com.Service
}

func NewWeatherService(service *weatherapi_com.Service) *WeatherService {
	return &WeatherService{service}
}

func (s *WeatherService) Forecast(ctx context.Context, city string, days int) ([]model.Weather, error) {
	forecast, err := s.service.Forecast(ctx, city, days)
	if err != nil {
		return nil, err
	}
	weathers, err := toWeathers(*forecast)
	if err != nil {
		return nil, err
	}
	return weathers, nil
}

func toWeathers(forecast weatherapi_com.Forecast) ([]model.Weather, error) {
	timezoneOffset, err := utils.GetTimezoneOffset(forecast.Location.TzId)
	if err != nil {
		return nil, err
	}

	var weathers []model.Weather
	for _, forecastDay := range forecast.Forecast.Forecastday {
		weather := forecastDayToWeather(forecastDay, forecast.Location.Country, forecast.Location.Name, forecast.Location.TzId, int64(timezoneOffset.Seconds()))
		for _, forecastHour := range forecastDay.Hour {
			hourWeather := forecastHourToWeather(forecastHour, forecast.Location.Country, forecast.Location.Name, forecast.Location.TzId, int64(timezoneOffset.Seconds()))
			weather.HourlyWeathers = append(weather.HourlyWeathers, hourWeather)
		}
		weathers = append(weathers, weather)
	}
	return weathers, nil
}

func forecastDayToWeather(forecastDay weatherapi_com.ForecastDay, country, city, timezone string, timezoneOffset int64) model.Weather {
	startTime := time.Unix(forecastDay.DateEpoch, 0)
	endTime := startTime.Add(time.Hour)
	return model.Weather{
		Condition:             forecastDay.Day.Condition.Text,
		Icon:                  fillImageSchema(forecastDay.Day.Condition.Icon),
		StartTime:             &startTime,
		EndTime:               &endTime,
		Country:               country,
		City:                  city,
		Timezone:              timezone,
		TimezoneOffsetSeconds: timezoneOffset,

		AvgTempC: forecastDay.Day.AvgtempC,
		MinTempC: forecastDay.Day.MintempC,
		MaxTempC: forecastDay.Day.MaxtempC,

		AvgWindKph: forecastDay.Day.MaxwindKph,
		MinWindKph: forecastDay.Day.MaxwindKph,
		MaxWindKph: forecastDay.Day.MaxwindKph,
	}
}

func forecastHourToWeather(hour weatherapi_com.ForecastHour, country, city, timezone string, timezoneOffset int64) model.Weather {
	startTime := time.Unix(hour.TimeEpoch, 0)
	endTime := startTime.Add(time.Hour)
	return model.Weather{
		Condition:             hour.Condition.Text,
		Icon:                  fillImageSchema(hour.Condition.Icon),
		StartTime:             &startTime,
		EndTime:               &endTime,
		Country:               country,
		City:                  city,
		Timezone:              timezone,
		TimezoneOffsetSeconds: timezoneOffset,

		AvgTempC: hour.TempC,
		MinTempC: hour.TempC,
		MaxTempC: hour.TempC,

		AvgWindKph: hour.WindKph,
		MinWindKph: hour.WindKph,
		MaxWindKph: hour.WindKph,
	}
}

func fillImageSchema(imageUrl string) string {
	if strings.HasPrefix(imageUrl, "//") {
		imageUrl = "https:" + imageUrl
	}
	return imageUrl
}

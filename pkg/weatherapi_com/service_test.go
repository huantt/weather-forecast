package weatherapi_com

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestForecast(t *testing.T) {
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		t.Skipf("Missing WEATHER_API_KEY")
	}
	service := NewService(key)
	forecast, err := service.Forecast(context.Background(), "London", 5)
	assert.NoError(t, err)
	assert.NotNil(t, forecast)
	assert.NotEmpty(t, forecast.Forecast.Forecastday)
}

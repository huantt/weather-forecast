package weatherapi

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/huantt/weather-forecast/pkg/weatherapi_com"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed testdata/forecast.json
var data []byte

func TestToWeathers(t *testing.T) {
	var forecast *weatherapi_com.Forecast
	err := json.Unmarshal(data, &forecast)
	if err != nil {
		panic(err)
	}

	weathers, err := toWeathers(*forecast)
	assert.NoError(t, err)
	assert.NotEmpty(t, weathers)
	for _, weather := range weathers {
		assert.NotEmpty(t, weather.Country)
		assert.NotEmpty(t, weather.City)
		assert.NotEmpty(t, weather.Condition)
		assert.NotEmpty(t, weather.Timezone)
		assert.NotEmpty(t, weather.Icon)
		assert.NotEmpty(t, weather.StartTime)
		assert.NotEmpty(t, weather.EndTime)
	}
	d, err := json.Marshal(weathers)
	fmt.Println(string(d))
}

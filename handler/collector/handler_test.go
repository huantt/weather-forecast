package collector

import (
	_ "embed"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"weather_forecast/model"
)

//go:embed testdata/weathers.json
var weathersData []byte

func TestGenerateReadme(t *testing.T) {
	var weathers []model.Weather
	err := json.Unmarshal(weathersData, &weathers)
	if err != nil {
		panic(err)
	}

	readme, err := generateReadme(weathers, "data/README.md.template")
	assert.NoError(t, err)
	assert.NotNil(t, readme)
	assert.NotEmpty(t, *readme)
}

package collector

import (
	_ "embed"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
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

	// Construct the path to data/test.txt relative to the test file
	readme, err := generateReadme(weathers, "../../template/README.md.template")
	assert.NoError(t, err)
	assert.NotNil(t, readme)
	assert.NotEmpty(t, *readme)
	os.WriteFile(".README.md", []byte(*readme), 0644)
}

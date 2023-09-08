package weatherapi_com

const APIEndpoint = "https://api.weatherapi.com"

type Forecast struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
	Forecast struct {
		Forecastday []ForecastDay `json:"forecastday"`
	} `json:"forecast"`
}

type ForecastDay struct {
	Date      string `json:"date"`
	DateEpoch int64  `json:"date_epoch"`
	Day       struct {
		MaxtempC          float64 `json:"maxtemp_c"`
		MaxtempF          float64 `json:"maxtemp_f"`
		MintempC          float64 `json:"mintemp_c"`
		MintempF          float64 `json:"mintemp_f"`
		AvgtempC          float64 `json:"avgtemp_c"`
		AvgtempF          float64 `json:"avgtemp_f"`
		MaxwindMph        float64 `json:"maxwind_mph"`
		MaxwindKph        float64 `json:"maxwind_kph"`
		TotalprecipMm     float64 `json:"totalprecip_mm"`
		TotalprecipIn     float64 `json:"totalprecip_in"`
		TotalsnowCm       float64 `json:"totalsnow_cm"`
		AvgvisKm          float64 `json:"avgvis_km"`
		AvgvisMiles       float64 `json:"avgvis_miles"`
		Avghumidity       float64 `json:"avghumidity"`
		DailyWillItRain   int     `json:"daily_will_it_rain"`
		DailyChanceOfRain int     `json:"daily_chance_of_rain"`
		DailyWillItSnow   int     `json:"daily_will_it_snow"`
		DailyChanceOfSnow int     `json:"daily_chance_of_snow"`
		Condition         struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		Uv float64 `json:"uv"`
	} `json:"day"`
	Astro struct {
		Sunrise          string `json:"sunrise"`
		Sunset           string `json:"sunset"`
		Moonrise         string `json:"moonrise"`
		Moonset          string `json:"moonset"`
		MoonPhase        string `json:"moon_phase"`
		MoonIllumination string `json:"moon_illumination"`
		IsMoonUp         int    `json:"is_moon_up"`
		IsSunUp          int    `json:"is_sun_up"`
	} `json:"astro"`
	Hour []ForecastHour `json:"hour"`
}

type ForecastHour struct {
	TimeEpoch int64   `json:"time_epoch"`
	Time      string  `json:"time"`
	TempC     float64 `json:"temp_c"`
	TempF     float64 `json:"temp_f"`
	IsDay     int     `json:"is_day"`
	Condition struct {
		Text string `json:"text"`
		Icon string `json:"icon"`
		Code int    `json:"code"`
	} `json:"condition"`
	WindMph      float64 `json:"wind_mph"`
	WindKph      float64 `json:"wind_kph"`
	WindDegree   int     `json:"wind_degree"`
	WindDir      string  `json:"wind_dir"`
	PressureMb   float64 `json:"pressure_mb"`
	PressureIn   float64 `json:"pressure_in"`
	PrecipMm     float64 `json:"precip_mm"`
	PrecipIn     float64 `json:"precip_in"`
	Humidity     int     `json:"humidity"`
	Cloud        int     `json:"cloud"`
	FeelslikeC   float64 `json:"feelslike_c"`
	FeelslikeF   float64 `json:"feelslike_f"`
	WindchillC   float64 `json:"windchill_c"`
	WindchillF   float64 `json:"windchill_f"`
	HeatindexC   float64 `json:"heatindex_c"`
	HeatindexF   float64 `json:"heatindex_f"`
	DewpointC    float64 `json:"dewpoint_c"`
	DewpointF    float64 `json:"dewpoint_f"`
	WillItRain   int     `json:"will_it_rain"`
	ChanceOfRain int     `json:"chance_of_rain"`
	WillItSnow   int     `json:"will_it_snow"`
	ChanceOfSnow int     `json:"chance_of_snow"`
	VisKm        float64 `json:"vis_km"`
	VisMiles     float64 `json:"vis_miles"`
	GustMph      float64 `json:"gust_mph"`
	GustKph      float64 `json:"gust_kph"`
	Uv           float64 `json:"uv"`
}

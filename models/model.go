package models

type Forecast struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Periods []Period `json:"periods"`
}

type Period struct {
	Number           int64  `json:"number"`
	Name             string `json:"name"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	IsDaytime        bool   `json:"isDaytime"`
	Temperature      int64  `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	TemperatureTrend string `json:"temperatureTrend"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
	ShortForecast    string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
}

// type ForecastHourly struct {
// 	Number        int    `json:"number"`
// 	StartTime     string `json:"start_time"`
// 	EndTime       string `json:"end_time"`
// 	IsDayTime     bool   `json:"is_day_time"`
// 	Temperature   int    `json:"temperature"`
// 	WindSpeed     string `json:"wind_speed"`
// 	WindDirection string `json:"wind_direction"`
// 	Icon          string `json:"icon"`
// 	ShortForecast string `json:"short_forecast"`
// }

// type TideData struct {
// 	Time  string `json:"time"`
// 	Value string `json:"value"`
// 	Hilo  string `json:"type"`
// }

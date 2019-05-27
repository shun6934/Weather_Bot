package lib

type WeatherResult struct {
	Dt   int `json:"dt"`
	Main struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Rain struct {
		OneH int `json:"1h"`
	} `json:"rain"`
	Snow struct {
		OneH int `json:"1h"`
	} `json:"snow"`
	Sys struct {
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Name     string `json:"name"`
	Coord    struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
}

type ForecastResult struct {
	List []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp     float64 `json:"temp"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
			Pressure float64 `json:"pressure"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Rain struct {
			OneH int `json:"1h"`
		} `json:"rain"`
		Snow struct {
			OneH int `json:"1h"`
		} `json:"snow"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
	} `json:"city"`
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type OpenWeatherMapConfig struct {
	OPEN_WEATHER_MAP_KEY string `required: "true" split_word: "true"`
	CITY_ID              string `required: "true" split_word: "true"`
}

type OpenWeatherMapAPIResponse struct {
	Main    Main      `json: "main"`
	Weather []Weather `json: "weather"`
	Coord   Coord     `json: "coord"`
	Wind    Wind      `json: "wind"`
	Dt      int64     `json: "dt"`
}

type Coord struct {
	Lon float64 `json: "lon"`
	Lat float64 `json: "lat"`
}

type Weather struct {
	Main        string `json: "main"`
	Description string `json: "description"`
	Icon        string `json: "icon"`
}

type Main struct {
	Temp      float64 `json: "temp"`
	FeelsLike float64 `json: "feels_like"`
	TempMin   float64 `json: "temp_min"`
	TempMax   float64 `json: "temp_max"`
	Pressure  int64   `json: "pressure"`
	Humidity  int64   `json: "humidity"`
}

type Wind struct {
	Speed float64 `json: "speed"`
	Deg   int64   `json: deg`
}

func main() {
	var owm OpenWeatherMapConfig
	if err := envconfig.Process("", &owm); err != nil {
		log.Fatalf("[Error] Failed load env var %s", err.Error())
	}

	endpoint := "https://api.openweathermap.org/data/2.5/weather"
	query := url.Values{}
	query.Set("APPID", owm.OPEN_WEATHER_MAP_KEY)
	query.Set("id", owm.CITY_ID)
	query.Set("lang", "ja")
	query.Set("units", "metric")

	res, err := http.Get(endpoint + "?" + query.Encode())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	var apiRes OpenWeatherMapAPIResponse
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}

	fmt.Printf("天気 %s", apiRes.Weather[0].Main)
	fmt.Printf("気温 %f", apiRes.Main.Temp)
	fmt.Printf("最高 %f", apiRes.Main.TempMax)
	fmt.Printf("最高 %f", apiRes.Main.TempMin)
}

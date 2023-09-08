package main

import (
  "fmt"
  "os"
  "net/http"
  "io"
  "time"
  "encoding/json"
)

type Weather struct {
  Location struct {
    Name string `json:"name"`
    Country string `json:""`
  } `json:"location"`
  Current struct {
    TempC float64 `json:"temp_c"`
    Condition struct {
	Text string `json:"text"`
    }`json:"condition"`
  } `json:"current"`
  Forecast struct {
    Forecastday []struct {
      Hour []struct {
        TimeEpoch int64 `json:"time_epoch"`
	TempC     float64 `json:"temp_c"`
	Condition struct {
	  Text string `json:"text"`
	} `json:"condition"`
	ChanceOfRain float64 `json:"chance_of_rain"`
      }
    } `json:"forecastday"`
  } `json:"forecast"`
}


func main() {
  q := "Samarkand"

  if len(os.Args) >= 2 {
    q = os.Args[1]
  }
  res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=<key>&q=" + q + "&days=1&alerts=no&aqi=no")
  if err != nil {
    panic(err)
  }

  defer res.Body.Close()

  if res.StatusCode != 200 {
    panic("API not working.")
  }

  body, err := io.ReadAll(res.Body)

  if err != nil {
    panic(err)
  }

  var weather Weather
  err = json.Unmarshal(body, &weather)
  if err != nil {
    panic(err)
  }
  fmt.Println("\n    Welcome to GoWeather (Powered by Ali)\n")
  location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

  fmt.Printf(
    "Joylashuv: %s\nHarorat: %.0fC\nHolat: %s\n\n",
    location.Name + "/" + location.Country,
    current.TempC,
    current.Condition.Text,
  )
  for _, hour := range hours {
    date := time.Unix(hour.TimeEpoch, 0)
    if date.Before(time.Now()) {
      continue
    }
    fmt.Printf(
      "%s - %.0fC, %.0f%%, %s\n",
      date.Format("15:00"),
      hour.TempC,
      hour.ChanceOfRain,
      hour.Condition.Text,
    )
  }
}

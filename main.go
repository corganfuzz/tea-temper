package main

import (
	"log"
	"os"
	"time"

	"github.com/tea-temper/climacell"
)

func main() {

	c := climacell.New(os.Getenv("CLIMACELL_API_KEY"))
	weatherSamples, err := c.HourlyForecast(climacell.ForecastArgs{
		LatLon:     &climacell.LatLon{Lat: 42.3826, Lon: -71.146},
		UnitSystem: "us",
		Fields:     []string{"temp"},
		StartTime:  time.Now(),
		EndTime:    time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Fatalf("error getting forecast data: %v", err)
	}

	var tempAtFive *climacell.FloatValue

	for i, w := range weatherSamples {
		if w.ObservationTime.Value.Hour() == 21 {
			tempAtFive = weatherSamples[i].Temp
			break
		}
	}

	if tempAtFive == nil || tempAtFive.Value == nil {
		log.Printf("No data on the temperature at 5, let's wing it! 🌺\n")
	} else if t := *tempAtFive.Value; t < 60 {
		log.Printf("It'll be %f out. Better make some hot tea! 🌺🍵\n", t)
	} else {
		log.Printf("It will be %f out. Iced tea it is! 🌺🍹\n", t)
	}

	// req, err := http.NewRequest(
	// 	http.MethodGet,
	// 	"https://api.climacell.co/v3/weather/forecast/hourly?lat=42.3826&lon=-71.1460&fields=temp",
	// 	nil,
	// )
	// if err != nil {
	// 	log.Fatalf("Error creating GET request: %v", err)
	// }

	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("apikey", os.Getenv("CLIMACELL_API_KEY"))

	// res, err := http.DefaultClient.Do(req)

	// if err != nil {
	// 	log.Fatalf("error sending HTTP request: %v", err)
	// }

	// responseBytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalf("error reding HTTP response body: %v", err)
	// }

	// var weatherSamples []climacell.Weather
	// if err := json.Unmarshal(responseBytes, &weatherSamples); err != nil {
	// 	log.Fatalf("error deserializing weather data")
	// }

	// for _, w := range weatherSamples {
	// 	if w.Temp != nil && w.Temp.Value != nil {
	// 		log.Printf("The temperature at %s is %f degrees %s\n",
	// 			w.ObservationTime.Value, *w.Temp.Value, w.Temp.Units)
	// 	} else {
	// 		log.Printf("No temperature data available at %s\n",
	// 			w.ObservationTime.Value)
	// 	}
	// }

	// if err != nil {
	// 	log.Fatalf("error reading HTTP response Body: %v", err)
	// }

	// log.Println("We got a response:", string(responseBytes))

}

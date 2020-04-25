package climacell

import (
	"net/http"
	"net/url"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.climacell.co",
	Path:   "/v3/",
}

//Client comment
type Client struct {
	c      *http.Client
	apiKey string
}

//LatLon properties
type LatLon struct { Lat, Lon float64}

//ForecastArgs properties
type ForecastArgs struct {
	LatLon 			*LatLon
	LocationID 		string
	UnitSystem 		string
	Fields 			[]string
	StartTime 		time.Time
	EndTime 		time.Time
}

// New comment
func New(apiKey string) *Client {
//	c := &http.Client{}
	c := &http.Client{Timeout: time.Minute}
	return &Client{
		c:      c,
		apiKey: apiKey,
	}
}

func (c *Client) HourlyForecast(queryParams) ([]Weather, error) {
	req, err := http.NewRequest("GET", /* URL */ , nil)
	if err != nil {
		return nil, err
	}

	// authenticate the request

	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", c.apiKey)
	
	// add queryparams to request

	defer res.Body.Close()
	var weatherSamples []Weather
	if err := json.NewDecoder(res.Body).Decode(&weatherSamples); err != nil {
		return nil, err
	}
	return weatherSamples, nil

	res, err := c.c.Do()
	if err != nil {
		return nil, err
	}

	// deserialize success or erro response and return its data
}

package climacell

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// FloatValue Comment
type FloatValue struct {
	Value *float64
	Units string
}

//NonNullableTimeValue Value Comment
type NonNullableTimeValue struct{ Value time.Time }

//Weather API comment
type Weather struct {
	Lat             float64              `json:"lat"`
	Lon             float64              `json:"lon"`
	Temp            *FloatValue          `json:"temp"`
	ObservationTime NonNullableTimeValue `json:"observation_time"`
}

func (args ForecastArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.LatLon != nil {
		q.Add("lat", strconv.FormatFloat(args.LatLon.Lat, 'f', -1, 64))
		q.Add("lon", strconv.FormatFloat(args.LatLon.Lon, 'f', -1, 64))
	}

	if args.LocationID != "" {
		q.Add("location_id", args.LocationID)
	}

	if args.UnitSystem != "" {
		q.Add("unit_system", args.UnitSystem)
	}

	if len(args.Fields) > 0 {
		q.Add("fields", strings.Join(args.Fields, ","))
	}

	if !args.StartTime.IsZero() {
		q.Add("start_time", args.StartTime.Format(time.RFC3339))
	}

	if !args.EndTime.IsZero() {
		q.Add("end_time", args.EndTime.Format(time.RFC3339))
	}
	return q
}

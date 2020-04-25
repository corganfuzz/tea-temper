package climacell

import (
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

package entities

import "time"

type Price struct {
	Value          int32  `json:"value"`
	FormattedValue string `json:"formatted_value"`
}

type Date struct {
	Value          time.Time `json:"value"`
	FormattedValue string    `json:"formatted_value"`
}

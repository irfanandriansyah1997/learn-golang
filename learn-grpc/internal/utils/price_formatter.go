package utils

import (
	"errors"
	"fmt"
	"learn-grpc/internal/entities"

	"github.com/dustin/go-humanize"
)

func FormattingPrice(price any) (*entities.Price, error) {
	switch val := price.(type) {
	case int:
		return &entities.Price{
			Value:          int32(val),
			FormattedValue: fmt.Sprintf("Rp %s", humanize.CommafWithDigits(float64(val), 0)),
		}, nil
	}

	return nil, errors.New("price is not integer")
}

package calculator

import (
	"learn-error/customerror"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, &customerror.NotFoundError{Message: "Value no valid"}
	}

	return nilai / pembagi, nil

}

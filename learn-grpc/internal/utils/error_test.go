package utils_test

import (
	"errors"
	"fmt"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedError = "Sample Error Message"

func TestNewNotFoundError(t *testing.T) {
	input := utils.NewNotFoundError(expectedError)
	assert.Equal(t, fmt.Sprintf("[Not Found]: %s", expectedError), input.Error())
}

func TestPanicIfError(t *testing.T) {
	defer func() {
		err, ok := recover().(error)

		if !ok {
			t.Errorf("This method should be throw error")
		}

		if err != nil && err.Error() != expectedError {
			t.Errorf("This method throwing incorrect error: %s", err.Error())
		}
	}()

	utils.PanicIfError(errors.New(expectedError))
}

func TestPanicIfErrorWithNilArgs(t *testing.T) {
	testCases := []entities.Fixture[func(err error), any]{
		{
			Title:         "testing method PanicIfError",
			Input:         utils.PanicIfError,
			ExpectedValue: nil,
		},
		{
			Title:         "testing method PanicIfNotFoundError",
			Input:         utils.PanicIfNotFoundError,
			ExpectedValue: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			defer func() {
				_, ok := recover().(error)

				if ok {
					t.Errorf("This method should be not throwing error")
				}
			}()

			testCase.Input(nil)
		})
	}
}

func TestPanicIfNotFoundError(t *testing.T) {
	defer func() {
		err, ok := recover().(error)

		if !ok {
			t.Errorf("This method should be throw error")
		}

		if err != nil && err.Error() != fmt.Sprintf("[Not Found]: %s", expectedError) {
			t.Errorf("This method throwing incorrect error: %s", err.Error())
		}
	}()

	utils.PanicIfNotFoundError(errors.New(expectedError))
}

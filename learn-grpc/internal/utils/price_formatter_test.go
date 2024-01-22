package utils_test

import (
	"learn-grpc/internal/entities"
	"learn-grpc/internal/utils"
	"testing"
)

func TestFormattingPrice(t *testing.T) {

	testCases := []entities.Fixture[any, string]{
		{
			Title:         "integer argument",
			Input:         10000,
			ExpectedValue: "Rp 10,000",
			IsError:       false,
		},
		{
			Title:         "integer argument 2",
			Input:         5000000,
			ExpectedValue: "Rp 5,000,000",
			IsError:       false,
		},
		{
			Title:         "float argument",
			Input:         10.5,
			ExpectedValue: "",
			IsError:       true,
		},
		{
			Title:         "string argument",
			Input:         "hello world",
			ExpectedValue: "",
			IsError:       true,
		},
		{
			Title:         "boolean argument",
			Input:         true,
			ExpectedValue: "",
			IsError:       true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result, err := utils.FormattingPrice(testCase.Input)

			// INFO: Expected error should be nil
			if !testCase.IsError && err != nil {
				t.Errorf("FormattingPrice() = returning error %s", err.Error())
				return
			}

			// INFO: Expected error should be defined
			if testCase.IsError == true && err == nil {
				t.Error("FormattingPrice() should be returning error")
				return
			}

			// INFO: Expected output from method is similar with attribute ExpectedValue
			if result.FormattedValue != testCase.ExpectedValue {
				t.Errorf("FormattingPrice().FormattedValue returning mismatch output %s with %s", testCase.ExpectedValue, result.FormattedValue)
			}
		})
	}
}

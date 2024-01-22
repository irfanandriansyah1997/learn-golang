package entities

type Fixture[input any, output any] struct {
	Title         string
	Input         input
	ExpectedValue output
	IsError       bool
}

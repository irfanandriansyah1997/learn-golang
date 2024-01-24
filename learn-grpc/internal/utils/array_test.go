package utils_test

import (
	"learn-grpc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Run("testing accumulate number", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := utils.Reduce(number, func(prev int, current int) int {
			return prev + current
		}, 0)

		assert.Equal(t, 55, result)
	})

	t.Run("testing filter data with status is true", func(t *testing.T) {
		type Args struct {
			value  int
			status bool
		}

		number := []Args{{10, true}, {5, true}, {6, false}, {3, false}}
		result := utils.Reduce(number, func(prev []Args, current Args) []Args {
			if current.status {
				prev = append(prev, current)
			}

			return prev
		}, make([]Args, 0))

		assert.Len(t, result, 2)
		assert.Equal(t, result[0].value, 10)
		assert.Equal(t, result[1].value, 5)
	})
}

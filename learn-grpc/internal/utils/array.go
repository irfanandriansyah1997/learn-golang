package utils

type ReducerFn[args any, output any] func(prev output, current args) output

func Reduce[args any, output any](arr []args, fn ReducerFn[args, output], init output) output {
	prev := init

	for _, item := range arr {
		prev = fn(prev, item)
	}

	return prev
}

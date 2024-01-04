package learncontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	contextD := context.WithValue(contextB, "d", "D")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextD)
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(time.Second * 1)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	ctx, cancel := context.WithCancel(context.Background())

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter: ", n)

		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(time.Second * 3)

	//  check goroutine jumlah nya yang jalan (goroutine leak)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter: ", n)
	}

	time.Sleep(time.Second * 1)

	//  check goroutine jumlah nya yang jalan (goroutine leak)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter: ", n)
	}

	time.Sleep(time.Second * 1)

	//  check goroutine jumlah nya yang jalan (goroutine leak)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

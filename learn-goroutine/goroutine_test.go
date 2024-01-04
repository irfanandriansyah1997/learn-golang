package main

import (
	"fmt"
	"testing"
	"time"
)

func runHelloWorld() {
	fmt.Println("Hello world")
}

func TestGoroutine(t *testing.T) {
	go runHelloWorld()

	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	fn := func(n int) {
		fmt.Println("Display ", n)
	}
	for i := 0; i < 100000; i++ {
		go fn(i)
	}

	time.Sleep(5 * time.Second)
}

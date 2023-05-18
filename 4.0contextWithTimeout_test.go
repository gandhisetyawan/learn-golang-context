package learn_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	//cek setelah jalankan, goroutine naik apa tidak
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)

	}

	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())

}

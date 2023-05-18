package learn_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// //goroutine leak
// func CreateCounter() chan int {
// 	destinatin := make(chan int)

// 	go func() {
// 		defer close(destinatin)
// 		counter := 1
// 		for {
// 			destinatin <- counter
// 			counter++
// 		}
// 	}()
// 	return destinatin
// }

// //goroutine leak
// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println(runtime.NumGoroutine())

// 	destinatin := CreateCounter()
// 	for n := range destinatin {
// 		fmt.Println("Counter", n)
// 		if n == 10 {
// 			break
// 		}
// 	}
// 	fmt.Println(runtime.NumGoroutine())

// }

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
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination

}
func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	//cek setelah jalankan goroutine naik apa tidak
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())

}

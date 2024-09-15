package belajargolanjutan

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// ada dua cara membuat context dengan background dan todo
	// keduanya sama" membuat context kosong
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestCreateContext(t *testing.T) {
	contextParent := context.Background()

	// value context itu bersifat `Pair-value`
	childA := context.WithValue(contextParent, "key1", "abdu")
	childB := context.WithValue(contextParent, "key2", "eunha")

	parentChildA := context.WithValue(childA, "childA1", "yerin")

	fmt.Println(childB)
	fmt.Println(parentChildA)
	fmt.Println(parentChildA.Value("childA1")) // yerin
	fmt.Println(parentChildA.Value("key2"))    // nil
}

// context with sinyal cancel
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
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())

	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)

	destination := CreateCounter(ctx)

	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context
	time.Sleep(2 * time.Second)
	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())
}

// cancel context with timeout
func TestContextWithCancelTimeout(t *testing.T) {
	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())

	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)

	destination := CreateCounter(ctx)

	defer cancel()

	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())
}

// context with deadline time
func TestContextWithCancelDeadline(t *testing.T) {
	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())

	parentCtx := context.Background()
	ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(10*time.Minute)) // cancel dlm waktu 10 mnt

	destination := CreateCounter(ctx)

	defer cancel()

	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Check total goroutine: ", runtime.NumGoroutine())
}

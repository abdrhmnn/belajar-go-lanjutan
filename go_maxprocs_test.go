package belajargolanjutan

import (
	"fmt"
	"runtime"
	"testing"
)

// ini untuk melihat jumlah cpu, thread serta goroutine yang sedang berjalan

func TestGomaxprocs(t *testing.T) {
	myCpu := runtime.NumCPU()
	fmt.Println("Total CPU yang dimiliki: ", myCpu)

	allThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Jumlah thread: ", allThread)

	allGoRoutine := runtime.NumGoroutine()
	fmt.Println("Jumlah goroutine yang sedang berjalan: ", allGoRoutine)
}

func TestEditGomaxprocs(t *testing.T) {
	myCpu := runtime.NumCPU()
	fmt.Println("Total CPU yang dimiliki: ", myCpu)

	// edit thread itu jarang dilakukan karean golang sudah memiliki sistem manajemen thread sendiri
	runtime.GOMAXPROCS(20) // edit jumlah thread

	allThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Jumlah thread: ", allThread)

	allGoRoutine := runtime.NumGoroutine()
	fmt.Println("Jumlah goroutine yang sedang berjalan: ", allGoRoutine)
}

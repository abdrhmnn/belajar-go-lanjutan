package belajargolanjutan

import (
	"fmt"
	"testing"
	"time"
)

// jika ada function yg ada return value nya dan berjalan dengan goroutine maka return value nya tidak bisa
// di tangkap kecuali menggunakan Channel

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

func TestRunHelloWorld(t *testing.T) {
	go RunHelloWorld() // menjalankan fungsi ini dengan goroutine atau berjalan secara async
	fmt.Println("Test!")

	// kadang" program bisa selesai terlebih dahulu tanpa menjalankan fungsi goroutine nya
	// untuk itu bisa diberikan sleep di program untuk menunggu proses goroutine selesai

	time.Sleep(1 * time.Second) // sleep for 1s
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		go RunHelloWorld()
	}

	time.Sleep(5 * time.Second)
}

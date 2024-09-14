package belajargolanjutan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// timer untuk menambahkan waktu di kode golang, dan juga bisa untuk mengirim suatu kejadian berdasarkan
// waktu yang sudah di setup

func TestTime(t *testing.T) {
	// ini cara manual untuk get hasil proses setelah waktu yang sudah ditentukan itu selesai
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("first: ", time.Now())

	time := <-timer.C // return nya adalah sebuah channel
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	// disini kita ga perlu setup timer nya jadi langsung ambil channel nya dengan fungsi `After`
	channel := time.After(5 * time.Second)
	fmt.Println("first: ", time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	// tanpa menggunakan channel jadi disini menggunakan function
	// AfterFunc jalan dengan goroutine jadi harus ditunggu dulu, bisa pakai time atau waitgroup
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println("first: ", time.Now())
	group.Wait()
}

// time.Ticker untuk representasi kejadian yang berulang
// jadi jika ingin memasukkan data ke channel berdasarkan waktu yang terjadi berkali-kali
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTickerWithTick(t *testing.T) {
	ticker := time.Tick(1 * time.Second)

	for tick := range ticker {
		fmt.Println(tick)
	}
}

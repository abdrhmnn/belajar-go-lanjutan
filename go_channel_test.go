package belajargolanjutan

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Channel adalah tempat komunikasi synchronous yang bisa dilakukan goroutine
// di channel ada yang namanya pengirim dan penerima, biasanya goroutine nya berbeda antara pengirim
// dan penerima

// saat mengirim data ke channel, maka goroutine akan ter-block sampai ada yang menerima data tersebut

// karateristik channel
// - by default hanya bisa menampung satu data, jika ingin menambah lagi harus menunggu data yang sudah
// berada di channel itu diambil
// - channel hanya bisa menerima satu jenis tipe data (atau bisa dibikin interface{})
// - channel bisa diambil lebih dari satu goroutine
// - channel harus di close jika sudah tidak digunakan lagi, biar tidak terjadi memory leak

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // create channel

	// mengirim data ke channel
	// channel <- "abdu"

	// menerima data dari channel
	// newData := <-channel
	// fmt.Println(<-channel) // menerima channel di parameter

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Abdu"
		fmt.Println("Berhasil insert data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	close(channel) // close channel
}

// channel as parameter
func ChannelWithParam(channel chan int) { // ini sudah langsung reference ke channel aslinya jadi gaperlu lagi
	// pakai pointer

	time.Sleep(2 * time.Second)
	channel <- 2_000_000
}

func TestChannelWithParam(t *testing.T) {
	channel := make(chan int)
	go ChannelWithParam(channel)

	fmt.Println(<-channel)
}

// channel in dan out
// bisa tentukan sebuah channel di param function itu hanya bisa mengirim atau menerima data channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Eunha"
}
func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println(data)
}

// buffered channel
// jadi saat mengirim data ke channel jika ingin goroutine nya tidak ter-block itu bisa dimasukkan dulu
// ke dalam buffered sih data" channel nya
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // capacity buffered 3

	// maka data yg sudah berada di dalam buffered channel itu gak harus menunggu dulu sampai data nya di ambil
	go func() {
		channel <- "abdu"
		channel <- "eunha"
		channel <- "sowon"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
}

// range channel, untuk menerima banyak data channel dengan pengulangan
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- strconv.Itoa(i)
		}

		// untuk menghentikan pengulangan saat data channel nya digunakan/loop
		// dan juga tidak terjadi error kelebihan capacity buffered
		close(channel)
	}()

	for data := range channel {
		fmt.Println("hasil ", data)
	}
}

// select channel, yaitu untuk menjalankan beberapa goroutine serta mendapatkan semua data dari channel tersebut
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go ChannelWithParam(channel1)
	go ChannelWithParam(channel2)

	// pengen pakai semua data dari channel yang sudah diisi
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		// default bisa ditambahkan untuk menunggu data dari channel yang di select itu sudah siap
		default:
			fmt.Println("Menunggu data channel")
		}

		if counter == 2 {
			break
		}
	}

	close(channel1)
	close(channel2)
}

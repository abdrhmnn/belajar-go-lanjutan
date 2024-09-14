package belajargolanjutan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// WaitGroup berfungsi untuk menunggu proses async itu selesai dilakukan

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // untuk mengurangi setiap nilai pada fungsi Add

	group.Add(1) // menambahkan 1 proses yang async seperti contohnya goroutine

	fmt.Println("Test")
	time.Sleep(1 * time.Second)
}

func TestRunAsync(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsynchronous(&group)
	}

	group.Wait()
	fmt.Println("Proses selesai!")
}

// sync.Once untuk menjalankan go routine hanya satu kali
var counter = 0

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			once.Do(func() {
				counter++
			})
		}()
		group.Done()
	}

	group.Wait()
	fmt.Println(counter)
}

// sync.Pool untuk menyimpan data dan menggunakannya sekaligus kembalikan data nya jika sudah tidak digunakan lagi
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "Test"
		},
	}

	pool.Put("Abdu")
	pool.Put("Eunha")
	pool.Put("Sowon")

	for i := 0; i <= 10; i++ {
		go func() {
			dataPool := pool.Get() // gunakan data yang tersimpan di dalam pool
			fmt.Println(dataPool)
			time.Sleep(1 * time.Second)
			pool.Put(dataPool) // balikin data lagi ke pool
		}()
	}

	time.Sleep(11 * time.Second)
}

// sync.Map, penggunaan map untuk goroutine
func AddMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	group := &sync.WaitGroup{}
	data := &sync.Map{}

	for i := 1; i <= 100; i++ {
		go AddMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, " : ", value)
		return true
	})
}

// sync.Cond, untuk membuat kondisi pada saat locking
var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i <= 10; i++ {
		go WaitCondition(i)
	}

	// pakai Signal, running satu persatu
	// go func() {
	// 	for i := 0; i <= 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	// langsung running sekaligus, pakai Broadcast
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Broadcast()
		}
	}()

	group.Wait()
}

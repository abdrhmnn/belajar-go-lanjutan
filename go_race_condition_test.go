package belajargolanjutan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// jadi nanti bisa saja bakal ada case dimana lebih dari satu goroutine itu mengakses variabel yang sama
// dan juga mengubah value nya, jika itu terjadi maka nanti akan menyebabkan `race condition`

// artinya bisa saja 2 goroutine itu mengubah variabel yang value nya itu tetap sama
// karena goroutine kan bisa jalan secara concurrent maupun parallel

func TestRaceCondition(t *testing.T) {
	variabel1 := 0

	// contohnya disini ada 1000 go routine yang mengakses variabel yang sama serta mengubah value nya juga
	for i := 1; i <= 1_000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				variabel1 += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(variabel1)
}

// untuk menghindari problem `race condition` itu bisa menggunakan Mutex (Mutual Exclusion)
// yaitu digunakan untuk melakukan locking dan unlocking, dimana data mutex yang sudah di locking itu
// tidak ada yang bisa melakukan locking lagi sampai dilakukannya unlocking
func TestRaceConditionWithMutex(t *testing.T) {
	variabel1 := 0
	var mutex sync.Mutex

	for i := 1; i <= 1_000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// jadi ini artinya hanya ada 1 goroutine yang bisa melakukan locking dan mengubah value variabel nya
				// jika proses nya sudah selesai itu di unlocking dan goroutine lain baru bisa melakukan locking kembali
				// ini seperti memberikan antrian kepada goroutine nya
				mutex.Lock()
				variabel1 += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(variabel1)
}

// RWMutex (Read and Write Mutex)
// misal ada case sebuah struct itu akan diakses oleh banyak goroutine
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	// write mutex
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	// read mutex
	account.RWMutex.Lock()
	balance := account.Balance
	account.RWMutex.Unlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			account.AddBalance(1)
			fmt.Println(account.GetBalance())
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println(account.GetBalance())
}

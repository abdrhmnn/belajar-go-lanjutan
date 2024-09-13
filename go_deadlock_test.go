package belajargolanjutan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// pada saat goroutine berjalan secara concurrenty ataupun parallel itu bisa saja terjadi yang namanya
// deadlock, yaitu proses goroutine yang menunggu lock sehingga tidak ada satupun goroutine yang berjalan

type UserBalance struct {
	Mutex   sync.Mutex
	Nama    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1 ", user1.Nama)
	user1.Change(amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2 ", user2.Nama)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Nama:    "Abdu",
		Balance: 100000,
	}

	user2 := UserBalance{
		Nama:    "Eunha",
		Balance: 100000,
	}

	// jadi ini kondisi nya, goroutine pertama akan lock user1 kemudian setelah itu ingin lock user2
	// nah sebelum goroutine pertama itu melakukan lock ke user2 ternyata user2 itu sudah keburuh di lock duluan
	// oleh goroutine yang kedua

	// nah dari case itu maka kedua goroutine ini akan terus menunggu locking sehingga tidak ada satupun
	// goroutine yang berjalan
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)
	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)
}

package helper

import (
	"fmt"
	"testing"
)

// aturan penamaan file unit testing nya yaitu cukup tambahkan `_test`
// dan didepan nama func unit testing nya itu ditambahkan `Test` serta harus ada parameter
// testing.T dan tidak mengembalikan return value

func TestSayHello(t *testing.T) {
	result := SayHello("abdu")
	if result != "Halo, abdu" {
		panic("Result is not 'Halo, abdu'")
	}
}

// menggagalkan test
// t.Fail() -> unit test dianggap gagal tetapi eksekusi unit test nya tetap dilakukan
// t.FailNow() -> unit test dianggap gagal tanpa melanjutkan eksekusi unit test nya
// t.Error() -> melakukan print log error dan juga sekaligus menjalankan t.Fail()
// t.Fatal() -> melakukan print log error dan juga sekaligus menjalankan t.FailNow()
func TestSayHello2(t *testing.T) {
	result := SayHello("eunha")
	if result != "Halo, abdu" {
		t.FailNow()
	}
}

func TestSayHello3(t *testing.T) {
	result := SayHello("eunha")
	if result != "Halo, abdu" {
		t.Fail()
	}

	fmt.Println("unit test tetap berjalan!")
}

func TestSayHello4(t *testing.T) {
	result := SayHello("eunha")
	if result != "Halo, abdu" {
		t.Error("Value bukan `Halo, abdu`")
	}

	fmt.Println("unit test with Error() tetap berjalan!")
}

func TestSayHello5(t *testing.T) {
	result := SayHello("eunha")
	if result != "Halo, abdu" {
		t.Fatal("Value bukan `Halo, abdu`")
	}

	fmt.Println("unit test with Fatal() tetap berjalan!")
}

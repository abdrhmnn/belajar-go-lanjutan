package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// testing assertion with Testify
func TestSayHello6(t *testing.T) {
	result := SayHello("eunha")
	assert.Equal(t, "Halo, abdu", result, "They should be equal!") // return with Fail()
	fmt.Println("Test with assertion is running!")
}

func TestSayHello7(t *testing.T) {
	result := SayHello("eunha")
	require.Equal(t, "Halo, abdu", result, "They should be equal!") // return with FailNow()
	fmt.Println("Test with require is running!")
}

// skip test
func TestSkipTest(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot do test on mac os")
	}

	fmt.Println("Test skip")
}

// before dan after test with TestMain
// harus buat func dengan nama TestMain kemudian paramnya testing.M
// func ini hanya akan di eksekusi sekali per go-lang package bukan tipa per-function unit test
func TestMain(m *testing.M) {
	// before
	fmt.Println("kondisi before fungsi unit test!")

	m.Run() // run all unit test in same package

	// after
	fmt.Println("kondisi after fungsi unit test!")
}

// sub test
func TestWithSubTest(t *testing.T) {
	t.Run("check for value abdu", func(t *testing.T) {
		result := SayHello("abdu")
		require.Equal(t, "Halo, abdu", result, "They should be equal!")
	})
	t.Run("check for value eunha", func(t *testing.T) {
		result := SayHello("eunha")
		require.Equal(t, "Halo, eunha", result, "They should be equal!")
	})
}

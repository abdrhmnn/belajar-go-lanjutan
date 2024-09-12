package benchmark

import (
	"testing"
)

// untuk test performa kecepatan kode
// setiap nama func benchmark itu harus diawali dengan Benchmark

// run benchmark
// go test -v -bench=.
// go test -v -run=<NamaUnitTestYangTidakAda> -bench=.
// go test -v -run=<NamaUnitTestYangTidakAda> -bench=<namaBenchmark>
// go test -v -bench=. ./...

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Eunha Sowon Umji Duar!!!!")
	}
}

// sub benchmark
func BenchmarkHelloWorldWithSub(b *testing.B) {
	b.Run("abdu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Eunha Sowon Umji Duar!!!!")
		}
	})

	b.Run("eunha", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Eunha Sowon Umji Duar!!!!")
		}
	})
}

// table test benchmark
func BenchmarkTableSayHello(t *testing.B) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Abdu",
			expected: "Halo, Abdu",
		},
		{
			name:     "eunha",
			expected: "Halo, Eunha",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.B) {
			SayHello(test.request)
		})
	}
}

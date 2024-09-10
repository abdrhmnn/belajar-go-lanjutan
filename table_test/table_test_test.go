package tabletest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// jika kita melakukan unit testing yang mirip2 seperti kondisi case testing nya, expected value nya, dll
// itu bisa dibikin dalam bentuk table test dengan bantuan data slice dari struct

func TestSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Abdu",
			request:  "Abdu",
			expected: "Hello, Abdu",
		},
		{
			name:     "eunha",
			request:  "Abdu",
			expected: "Hello, Eunha",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			require.Equal(t, test.expected, result, "They should be equal!")
		})
	}
}

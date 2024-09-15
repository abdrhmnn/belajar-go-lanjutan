package belajargolanjutan

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncodingJson(t *testing.T) {
	LogJson("Abdu")
	LogJson(1)
	LogJson(false)
	LogJson([]int{100, 200_000, 150322})
}

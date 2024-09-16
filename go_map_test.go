package belajargolanjutan

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Map object
func TestEncodeMapObject(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"nama":  "HP Samsung",
		"price": 2_000_000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestDecodeMapObject(t *testing.T) {
	product := `{"id":"P0001","nama":"HP Samsung","price":2000000}`
	var result map[string]interface{}

	err := json.Unmarshal([]byte(product), &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["nama"])
	fmt.Println(result["price"])
}

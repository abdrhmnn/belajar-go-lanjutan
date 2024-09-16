package belajargolanjutan

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Biodata struct {
	FirstName  string
	MiddleName string
	LastName   string
}

// decode streaming data (bisa berupa file, I/O Reader, body dari respon web, dll)
func TestDecodeStreaming(t *testing.T) {
	reader, _ := os.Open("testing.json")
	decoder := json.NewDecoder(reader)

	biodata := &Biodata{}
	decoder.Decode(biodata)

	fmt.Println(biodata)
}

// encode streaming
func TestEncodeStreaming(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	biodata := &Biodata{
		FirstName:  "abdu",
		MiddleName: "abdu_rahman",
		LastName:   "abdu3",
	}

	encoder.Encode(biodata)
	fmt.Println("selesai melakukan encode!")
}

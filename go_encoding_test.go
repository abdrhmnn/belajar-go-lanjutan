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

// json object
type DataObject struct {
	Nama   string
	Umur   int
	Kota   []string
	Status bool
}

func TestJsonObject(t *testing.T) {
	dataObject := DataObject{
		Nama:   "Abdu",
		Umur:   24,
		Kota:   []string{"Jakarta", "Bandung", "Medan"},
		Status: false,
	}

	bytes, _ := json.Marshal(dataObject)
	fmt.Println(string(bytes))
}

// decode json
func TestDecodeJson(t *testing.T) {
	// konversi ke array byte dulu terus tinggal di mapping ke struct
	dataJson := `{"Nama":"Abdu","Umur":24,"Kota":["Jakarta","Bandung","Medan"],"Status":false}`
	konversiBytes := []byte(dataJson)

	dataObject := &DataObject{}

	err := json.Unmarshal(konversiBytes, dataObject)

	if err != nil {
		panic(err)
	}

	fmt.Println(dataObject)
}

// JSON Array with slice
type DataObject2 struct {
	Nama   string
	Umur   int
	Alamat string
}

type ParentDataObject2 struct {
	Lokasi   string
	Hobby    []string
	DataDiri []DataObject2
}

func TestSliceArray(t *testing.T) {
	parentDataObject2 := ParentDataObject2{
		Lokasi: "Jakarta",
		Hobby:  []string{"Game", "Tidur"},
		DataDiri: []DataObject2{
			{
				Nama:   "Abdu",
				Umur:   24,
				Alamat: "jkt",
			},
			{
				Nama:   "Eunha",
				Umur:   24,
				Alamat: "bdg",
			},
		},
	}

	bytes, _ := json.Marshal(parentDataObject2)
	fmt.Println(string(bytes))
	fmt.Println(parentDataObject2.Lokasi)
}

func TestArrayJson(t *testing.T) {
	dataDiri := []DataObject2{
		{
			Nama:   "Abdu",
			Umur:   24,
			Alamat: "jkt",
		},
		{
			Nama:   "Eunha",
			Umur:   24,
			Alamat: "bdg",
		},
	}

	bytes, _ := json.Marshal(dataDiri)
	fmt.Println(string(bytes))
}

func TestArrayObject(t *testing.T) {
	arrayObject := `{"Lokasi":"Jakarta","Hobby":["Game","Tidur"],"DataDiri":[{"Nama":"Abdu","Umur":24,"Alamat":"jkt"},{"Nama":"Eunha","Umur":24,"Alamat":"bdg"}]}`

	arraySliceObject := `[{"Nama":"Abdu","Umur":24,"Alamat":"jkt"},{"Nama":"Eunha","Umur":24,"Alamat":"bdg"}]`

	konversiBytes := []byte(arrayObject)
	konversiBytesSlice := []byte(arraySliceObject)

	parentObject := &ParentDataObject2{}
	parentSliceObject := &[]DataObject2{} // slice data

	err := json.Unmarshal(konversiBytes, parentObject)
	errSlice := json.Unmarshal(konversiBytesSlice, parentSliceObject)

	if err != nil || errSlice != nil {
		panic(err)
	}

	fmt.Println(parentObject)
	fmt.Println(parentObject.Lokasi)
	fmt.Println(parentObject.DataDiri[0:1])
	fmt.Println(parentSliceObject)
}

// Tag Json
type DataObject3 struct {
	Nama   string `json:"nama_saya"`
	Alamat string `json:"alamat_saya"`
	Umur   int    `json:"umur_saya"`
}

func TestWriteTagJson(t *testing.T) {
	dataObject3 := DataObject3{
		Nama:   "abdu",
		Alamat: "jkt",
		Umur:   24,
	}

	bytes, _ := json.Marshal(dataObject3)
	fmt.Println(string(bytes))
}

func TestReadTagJson(t *testing.T) {
	// untuk read go-lang tidak memperdulikan penamaan key nya (mau huruf kecil atau besar)
	newArr := `{"nama_saya":"abdu","alamat_saya":"jkt","umur_saya":24}`
	dataObject3 := &DataObject3{}
	err := json.Unmarshal([]byte(newArr), dataObject3)

	if err != nil {
		panic(err)
	}

	fmt.Println(dataObject3)
}

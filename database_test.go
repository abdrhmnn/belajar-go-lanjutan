package belajargolanjutan

import (
	"belajar-golang-database/db"
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestKonekDb(t *testing.T) {
	database := db.GetConnection()
	defer database.Close()

	fmt.Println("Connected!")
}

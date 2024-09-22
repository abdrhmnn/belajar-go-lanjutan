package belajargolanjutan

import (
	"belajar-golang-database/db"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestInsertData(t *testing.T) {
	database := db.GetConnection()
	defer database.Close()

	ctx := context.Background()

	// script := "INSERT INTO users(name, age, email, rating, birth_date, married) VALUES('eunha', 20, 'testing', 10.25, '2000-12-2', false)"
	script := "INSERT INTO users(name, age, email, rating, birth_date, married) VALUES('sowon', 20, NULL, 10.25, '2000-12-2', false)"
	_, err := database.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil ditambahkan!")
}

type USERS struct {
	ID     int
	NAMA   string
	AGE    int
	EMAIL  sql.NullString // untuk handling column yang bisa nullable
	RATING sql.NullFloat64
	// setiap data di db yg tipe DATE, DATETIME, TIME, TIMESTAMP, itu representasi di golang menjadi time.Time
	// ketika query data time.Time itu akan return di golang menjadi []byte atau []uint8
	// maka harus konversi ke bentuk string dulu
	// kalo mau otomatis dilakukan oleh mysql-driver itu tinggal tambahkan `parseTime=true` di DNS nya
	CREATED_AT sql.NullTime
	BIRTH_DATE sql.NullTime
	MARRIED    sql.NullBool
}

func TestReadData(t *testing.T) {
	database := db.GetConnection()
	defer database.Close()

	ctx := context.Background()

	script := "SELECT * FROM users"
	rows, err := database.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var users USERS
		err := rows.Scan(&users.ID, &users.NAMA, &users.AGE, &users.EMAIL, &users.RATING, &users.CREATED_AT, &users.BIRTH_DATE, &users.MARRIED)
		if err != nil {
			fmt.Println("Kesalahan pada read data!")
		}

		fmt.Println(users.ID)
		fmt.Println(users.NAMA)
		fmt.Println(users.AGE)
		fmt.Println(users.EMAIL)
		fmt.Println(users.RATING)
		fmt.Println(users.CREATED_AT)
		fmt.Println(users.BIRTH_DATE)
		fmt.Println(users.MARRIED)
	}
}

// query sql with parameter
func TestQueryWithParam(t *testing.T) {
	eunha := "yerin"
	age := 30
	email := "testing@gmail.com"
	rating := 100.53
	birth_date := "2000-10-23"
	married := true

	database := db.GetConnection()
	defer database.Close()

	ctx := context.Background()

	script := "INSERT INTO users(name, age, email, rating, birth_date, married) VALUES(?,?,?,?,?,?)"

	// bisa juga pakai QueryContext
	result, err := database.ExecContext(ctx, script, eunha, age, email, rating, birth_date, married)
	if err != nil {
		panic(err)
	}

	data, _ := result.LastInsertId()
	fmt.Println("Data telah ditambahkan dengan id: ", data)

	fmt.Println("Data berhasil ditambahkan!")
}

// query with prepare statement
func TestQueryWithPrepareStatement(t *testing.T) {
	database := db.GetConnection()
	defer database.Close()

	ctx := context.Background()

	script := "INSERT INTO users(name, age, email, rating, birth_date, married) VALUES(?,?,?,?,?,?)"

	// perbedaan menggunakan prepare statement dan menggunakan query atau exec with parameter yaitu
	// kalo pakai query atau exec itu sebenarnya membuat prepare statement juga tapi untuk 1 query
	// jadi kalo kita melakukan 100 query berarti kita sama aja membuat 100 prepare statement
	// dimana 1 prepare statment itu sudah include 1 koneksi db jadi sama halnya membuat 100 koneksi db
	// nah ini akan mempengaruhi database pooling nya

	// kalo pakai prepare statement langsung/manual itu cukup bikin 1 koneksi db yang bisa running
	// banyak query sekaligus
	statement, err := database.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	// test create data banyak dalam satu kali koneksi
	for i := 1; i <= 10; i++ {
		eunha := "yerin"
		age := 30
		email := "testing@gmail.com"
		rating := 100.53
		birth_date := "2000-10-23"
		married := true

		time.Sleep(1 * time.Second)
		result, _ := statement.ExecContext(ctx, eunha, age, email, rating, birth_date, married)
		id, _ := result.LastInsertId()

		fmt.Println("Data id: ", id)
	}

	fmt.Println("Data berhasil ditambahkan!")
}

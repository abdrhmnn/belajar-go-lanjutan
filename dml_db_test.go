package belajargolanjutan

import (
	"belajar-golang-database/db"
	"context"
	"fmt"
	"testing"
)

func TestInsertData(t *testing.T) {
	database := db.GetConnection()
	defer database.Close()

	ctx := context.Background()

	script := "INSERT INTO users(name, age, email) VALUES('eunha', 20, 'testing')"
	_, err := database.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil ditambahkan!")
}

type USERS struct {
	ID    int
	NAMA  string
	AGE   int
	EMAIL string
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
		err := rows.Scan(&users.ID, &users.NAMA, &users.AGE, &users.EMAIL)
		if err != nil {
			fmt.Println("Kesalahan pada read data!")
		}

		fmt.Println(users.ID)
		fmt.Println(users.NAMA)
		fmt.Println(users.AGE)
		fmt.Println(users.EMAIL)
	}
}

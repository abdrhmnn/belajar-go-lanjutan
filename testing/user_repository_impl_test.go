package testing

import (
	"belajar-golang-database/db"
	"belajar-golang-database/entity"
	repositoryimpl "belajar-golang-database/repository_impl"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestInsertData(t *testing.T) {
	userRepository := repositoryimpl.NewUserRepository(db.GetConnection())

	ctx := context.Background()
	user := entity.USERS{
		NAMA:       "sowon",
		AGE:        200,
		EMAIL:      sql.NullString{String: "sowon@gmail.com", Valid: true},
		RATING:     sql.NullFloat64{Float64: 200.124, Valid: true},
		BIRTH_DATE: sql.NullTime{Time: time.Now(), Valid: true},
		MARRIED:    sql.NullBool{Bool: false, Valid: true},
	}

	_, err := userRepository.Insert(ctx, user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil insert data!")
}

func TestFindById(t *testing.T) {
	userRepository := repositoryimpl.NewUserRepository(db.GetConnection())

	ctx := context.Background()
	result, err := userRepository.FindById(ctx, 30)
	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil menemukan data!")
	fmt.Println("nama: ", result.NAMA)
	fmt.Println("umur: ", result.AGE)
	fmt.Println("email: ", result.EMAIL.String)
}

func TestFindAll(t *testing.T) {
	userRepository := repositoryimpl.NewUserRepository(db.GetConnection())

	ctx := context.Background()
	result, err := userRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, hasil := range result {
		fmt.Println(hasil)
	}
}

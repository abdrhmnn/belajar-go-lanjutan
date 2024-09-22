package repositoryimpl

import (
	"belajar-golang-database/entity"
	"belajar-golang-database/repository"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type repositoryUserImpl struct {
	DB *sql.DB
}

// method untuk implementasi repository nya
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &repositoryUserImpl{DB: db}
}

func (repository *repositoryUserImpl) Insert(ctx context.Context, user entity.USERS) (entity.USERS, error) {
	script := "INSERT INTO users(name, age, email, rating, birth_date, married) VALUES(?,?,?,?,?,?)"
	result, err := repository.DB.ExecContext(ctx, script, user.NAMA, user.AGE, user.EMAIL, user.RATING, user.BIRTH_DATE, user.MARRIED)
	if err != nil {
		return user, err
	}

	id, _ := result.LastInsertId()
	user.ID = int32(id)
	return user, nil
}

func (repository *repositoryUserImpl) FindById(ctx context.Context, id int32) (entity.USERS, error) {
	script := "SELECT name, age, email, rating, birth_date, married FROM users WHERE id = ? LIMIT 1"
	result, err := repository.DB.QueryContext(ctx, script, id)
	user := entity.USERS{}
	if err != nil {
		return user, err
	}
	defer result.Close()

	if result.Next() {
		result.Scan(&user.NAMA, &user.AGE, &user.EMAIL, &user.RATING, &user.BIRTH_DATE, &user.MARRIED)
		return user, nil
	} else {
		return user, errors.New("Id " + strconv.Itoa(int(id)) + " tidak ditemukan!")
	}
}

func (repository *repositoryUserImpl) FindAll(ctx context.Context) ([]entity.USERS, error) {
	script := "SELECT name, age, email, rating, birth_date, married FROM users"
	result, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var users []entity.USERS
	for result.Next() {
		user := entity.USERS{}
		result.Scan(&user.NAMA, &user.AGE, &user.EMAIL, &user.RATING, &user.BIRTH_DATE, &user.MARRIED)
		users = append(users, user)
	}

	return users, nil
}

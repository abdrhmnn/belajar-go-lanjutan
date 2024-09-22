package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.USERS) (entity.USERS, error)
	FindById(ctx context.Context, id int32) (entity.USERS, error)
	FindAll(ctx context.Context) ([]entity.USERS, error)
}

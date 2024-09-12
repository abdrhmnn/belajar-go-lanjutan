package service

import (
	"belajar-golang-unit-test/mocking"
	entity "belajar-golang-unit-test/mocking/Entity"
	"errors"
)

type CategoryService struct {
	Repository mocking.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found!")
	} else {
		return category, nil
	}
}

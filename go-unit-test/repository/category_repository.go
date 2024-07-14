package repository

import "belajar-unit-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

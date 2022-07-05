package repo

import "gin_latihan/model"

type CategoryRepository interface {
	Add(newCategory *model.Category) []model.Category
}

type categoryRepository struct {
	categoryDb []model.Category
}

func (c *categoryRepository) Add(newCategory *model.Category) []model.Category {
	c.categoryDb = append(c.categoryDb, *newCategory)
	return c.categoryDb
}

func NewCategoryRepository() CategoryRepository {
	repo := new(categoryRepository)
	return repo
}

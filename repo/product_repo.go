package repo

import "gin_latihan/model"

type ProductRepository interface {
	Add(newProduct *model.Product) []model.Product
}

type productRepository struct {
	productDb []model.Product
}

func (p *productRepository) Add(newProduct *model.Product) []model.Product {
	p.productDb = append(p.productDb, *newProduct)
	return p.productDb
}

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}

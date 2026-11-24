package repository

import (
	"Mini_Store_Management/model"
	"errors"
	"math"
	"slices"
)

type Repository struct {
	Products []model.Product
}

func NewRepository() Repository {
	return Repository{
		Products: make([]model.Product, 0),
	}
}

func (r *Repository) Add(p model.Product) error {
	err := p.Validate()
	if err != nil {
		return err
	}

	i := r.FindProductIndex(p.Id())
	if i != -1 {
		return errors.New("product already exists")
	}
	r.Products = append(r.Products, p)
	return nil
}

func (r *Repository) FindProductIndex(id int) int {
	return slices.IndexFunc(r.Products, func(product model.Product) bool {
		return product.Id() == id
	})
}

func (r *Repository) FindByID(id int) (model.Product, error) {
	index := r.FindProductIndex(id)
	if index == -1 {
		return model.Product{}, errors.New("product not found")
	}
	product := r.Products[index]

	return product, nil
}

func (r *Repository) IncreaseStock(id int, quantity int) error {
	i := r.FindProductIndex(id)
	if i == -1 {
		return errors.New("product not found")
	}

	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	if quantity > math.MaxInt-r.Products[i].Stock() {
		return errors.New("quantity is so great")
	}

	r.Products[i].SetStock(r.Products[i].Stock() + quantity)
	return nil
}

func (r *Repository) DecreaseStock(id int, quantity int) error {
	i := r.FindProductIndex(id)
	if i == -1 {
		return errors.New("product not found")
	}

	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	if quantity > r.Products[i].Stock() {
		return errors.New("quantity is greater than the stock")
	}

	r.Products[i].SetStock(r.Products[i].Stock() - quantity)
	return nil
}

func (r *Repository) Purchase(id int, quantity int) (int, error) {
	i := r.FindProductIndex(id)
	if i == -1 {
		return 0, errors.New("product not found")
	}

	if quantity <= 0 {
		return 0, errors.New("quantity must be positive")
	}

	product := r.Products[i]
	if quantity > product.Stock() {
		return 0, errors.New("quantity is greater than the stock")
	}

	if quantity > math.MaxInt/product.Price() {
		return 0, errors.New("total price overflow")
	}

	r.Products[i].SetStock(r.Products[i].Stock() - quantity)
	return quantity * product.Price(), nil
}

func (r *Repository) Remove(id int) error {
	i := r.FindProductIndex(id)
	if i == -1 {
		return errors.New("product not found")
	}

	if r.Products[i].Stock() != 0 {
		return errors.New("product stock is not zero")
	}

	r.Products = slices.Delete(r.Products, i, i+1)
	return nil
}

package repository

import (
	"Mini_Store_Management/model"
	"errors"
	"math"
	"slices"
)

type Repository struct {
	products []model.Product
}

func NewRepository() Repository {
	return Repository{
		products: make([]model.Product, 0),
	}
}

func (r *Repository) ListProducts() []model.Product {
	p := make([]model.Product, len(r.products))
	copy(p, r.products)
	return p
}

func (r *Repository) Add(p model.Product) error {
	err := p.Validate()
	if err != nil {
		return err
	}

	i := r.FindProductIndex(p.ID())
	if i != -1 {
		return errors.New("product already exists")
	}
	r.products = append(r.products, p)
	return nil
}

func (r *Repository) FindProductIndex(id int) int {
	return slices.IndexFunc(r.products, func(product model.Product) bool {
		return product.ID() == id
	})
}

func (r *Repository) FindByID(id int) (model.Product, error) {
	index := r.FindProductIndex(id)
	if index == -1 {
		return model.Product{}, errors.New("product not found")
	}
	product := r.products[index]

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

	if quantity > math.MaxInt-r.products[i].Stock() {
		return errors.New("quantity is so great")
	}

	ok := r.products[i].SetStock(r.products[i].Stock() + quantity)
	if !ok {
		return errors.New("change stock failed")
	}
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

	if quantity > r.products[i].Stock() {
		return errors.New("quantity is greater than the stock")
	}

	ok := r.products[i].SetStock(r.products[i].Stock() - quantity)
	if !ok {
		return errors.New("change stock failed")
	}
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

	product := r.products[i]
	if quantity > product.Stock() {
		return 0, errors.New("quantity is greater than the stock")
	}

	if quantity > math.MaxInt/product.Price() {
		return 0, errors.New("total price overflow")
	}

	ok := r.products[i].SetStock(r.products[i].Stock() - quantity)
	if !ok {
		return 0, errors.New("change stock failed")
	}
	return quantity * product.Price(), nil
}

func (r *Repository) Remove(id int) error {
	i := r.FindProductIndex(id)
	if i == -1 {
		return errors.New("product not found")
	}

	if r.products[i].Stock() != 0 {
		return errors.New("product stock is not zero")
	}

	r.products = slices.Delete(r.products, i, i+1)
	return nil
}

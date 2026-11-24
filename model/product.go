package model

import (
	"errors"
	"strings"
)

type Product struct {
	id    int
	name  string
	price int
	stock int
}

func (p *Product) Price() int {
	return p.price
}

func (p *Product) SetPrice(price int) {
	p.price = price
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) Id() int {
	return p.id
}

func (p *Product) SetId(id int) {
	p.id = id
}

func (p *Product) Stock() int {
	return p.stock
}

func (p *Product) SetStock(stock int) {
	p.stock = stock
}

func NewProduct(id int, name string, price int, stock int) (Product, error) {
	product := Product{
		name:  name,
		price: price,
		stock: stock,
		id:    id,
	}
	err := product.Validate()
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.id <= 0 {
		return errors.New("invalid product id")
	}
	if len(strings.TrimSpace(p.name)) == 0 {
		return errors.New("name is empty")
	}
	if p.price <= 0 {
		return errors.New("price is less than or equal to 0")
	}
	if p.stock < 0 {
		return errors.New("stock is less than 0")
	}
	return nil
}

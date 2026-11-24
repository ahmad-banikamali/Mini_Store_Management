package repository

import (
	"Mini_Store_Management/model"
	"testing"
)

func TestListProducts(t *testing.T) {
	r := NewRepository()
	products := r.ListProducts()
	if len(products) != 0 {
		t.Errorf("ListProducts returned %d products, want 0", len(products))
	}
	p, err := model.NewProduct(10, "a", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(p)
	if err != nil {
		t.Fatal(err)
	}

	products = r.ListProducts()
	if len(products) != 1 {
		t.Errorf("ListProducts returned %d products, want 1", len(products))
	}

	product, err := model.NewProduct(1, "b", 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	products[0] = product

	got, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}

	if got != p {
		t.Error("replacing product in returned slice affected repository")
	}

	products[0].SetStock(2)

	got, err = r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}

	if got.Stock() != 1 {
		t.Errorf("repository stock = %d, want 1", got.Stock())
	}
}

func TestAddProduct(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	if len(r.products) != 1 {
		t.Error("Products length should be 1")
	}
}

func TestAddDuplicateProduct(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	err = r.Add(product)
	if err == nil {
		t.Error("expected error when adding duplicate product, got nil")
	}

	if len(r.products) != 1 {
		t.Error("Products length should be 1")
	}
}

func TestFindByID(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}

	id, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}

	if id != product {
		t.Error("Product does not match")
	}

}

func TestFindByIDNotFound(t *testing.T) {
	r := NewRepository()
	_, err := r.FindByID(10)
	if err == nil {
		t.Error("expected error")
	}
}

func TestIncreaseStock(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	err = r.IncreaseStock(10, 2)
	if err != nil {
		t.Fatal("Error increasing stock", err)
	}
	id, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}

	if id.Stock() != 3 {
		t.Error("Product does not match")
	}
}

func TestIncreaseStockInvalidQuantity(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	err = r.IncreaseStock(10, -2)
	if err == nil {
		t.Error("expected error")
	}
	id, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}
	if id.Stock() != 1 {
		t.Error("Product does not match")
	}
}

func TestDecreaseStock(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	err = r.DecreaseStock(10, 1)
	if err != nil {
		t.Fatal("Error decreasing stock", err)
	}

	id, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}

	if id.Stock() != 1 {
		t.Error("Product does not match")
	}
}

func TestDecreaseStockInsufficientStock(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	err = r.DecreaseStock(10, 3)
	if err == nil {
		t.Error("expected error")
	}
	id, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}
	if id.Stock() != 2 {
		t.Error("Product does not match")
	}
}

func TestPurchase(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 25, 4)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	purchase, err := r.Purchase(10, 2)
	if err != nil {
		t.Fatal(err)
	}
	if purchase != 50 {
		t.Error("Product does not match")
	}
	id, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}
	if id.Stock() != 2 {
		t.Error("Product does not match")
	}
}

func TestPurchaseInsufficientStock(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 25, 4)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	total, err := r.Purchase(10, 5)
	if err == nil {
		t.Error("expected error")
	}

	if total != 0 {
		t.Errorf("expected total 0, got %d", total)
	}

	id, err := r.FindByID(10)
	if err != nil {
		t.Fatal(err)
	}
	if id.Stock() != 4 {
		t.Error("Product does not match")
	}
}

func TestRemoveProduct(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 25, 0)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}

	err = r.Remove(10)
	if err != nil {
		t.Fatal("Error removing product", err)
	}
	_, err = r.FindByID(10)
	if err == nil {
		t.Fatal("expected error, product was still found")
	}
	if len(r.products) != 0 {
		t.Error("Products length should be 0")
	}
}

func TestRemoveProductInvalidProduct(t *testing.T) {
	r := NewRepository()
	product, err := model.NewProduct(10, "a", 25, 2)
	if err != nil {
		t.Fatal(err)
	}
	err = r.Add(product)
	if err != nil {
		t.Fatal("Error adding product", err)
	}
	err = r.Remove(10)
	if err == nil {
		t.Error("expected error")
	}
	if len(r.products) != 1 {
		t.Error("Products length should be 1")
	}
}

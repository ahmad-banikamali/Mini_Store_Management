package main

import (
	"Mini_Store_Management/model"
	"Mini_Store_Management/repository"
	"fmt"
)

func main() {
	p, e := model.NewProduct(10, "A", 1, 2)
	if e != nil {
		fmt.Printf("Error creating product: %s\n", e.Error())
		return
	}
	r := repository.NewRepository()
	err := r.Add(p)
	if err != nil {
		fmt.Println(err)
	}
	id, e := r.FindByID(10)
	if e != nil {
		fmt.Printf("Error finding product: %s\n", e.Error())
	}

	e = r.Remove(10)
	if e != nil {
		fmt.Printf("Error removing product: %s\n", e.Error())
	}

	fmt.Println(id)
}

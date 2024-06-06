package main

import (
	"fmt"
	"strconv"
)

type Order struct {
	Id        int
	Product   Product
	Customer  Customer
	Quantity  int
	Price     float32
	CreatedAt string
}

func (o Order) toString() {
	fmt.Printf("Id: %d\nTitre du produit: %d\nNom du client: %d\nQuantité: %d\nPrix: %.2f\nCreatedAt: %s\n", o.Id, o.Product.Title, o.Customer.LastName, o.Quantity, o.Price, o.CreatedAt)
}

func (o Order) toCSVRecord() []string {
	return []string{
		strconv.Itoa(o.Id),
		o.Customer.FirstName,
		o.Product.Title,
		strconv.Itoa(o.Quantity),
		fmt.Sprintf("%.2f", o.Price),
		o.CreatedAt,
	}
}

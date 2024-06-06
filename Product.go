package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Quantity    int
	Price       float32
	IsActive    bool
}

func (p Product) toString() {
	fmt.Printf("Title: %s, Description: %s, Quantity: %d, Price: %.2f, IsActive: %t \n", p.Title, p.Description, p.Quantity, p.Price, p.IsActive)
}
func (p Product) displayMostExpensive() {
	fmt.Printf("Le produit le plus cher est le %s, son prix est de %.2f € et il en reste %d unités \n", p.Title, p.Price, p.Quantity)
}

func (p Product) toCSVRecord() []string {
	return []string{
		p.Title,
		p.Description,
		strconv.Itoa(p.Quantity),
		fmt.Sprintf("%.2f", p.Price),
		strconv.FormatBool(p.IsActive), // Convert boolean to string
	}
}

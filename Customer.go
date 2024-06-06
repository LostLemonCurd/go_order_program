package main

import (
	"fmt"
)

type Customer struct {
	Id          int
	LastName    string
	FirstName   string
	PhoneNumber string
	Email       string
	Address     string
	TotalOrders float32
}

func (c Customer) toString() {
	fmt.Printf("Id: %d\nNom: %s\n, Prénom: %s\n, Numéro de Téléphone: %s\n, Email: %s\n, Adresse: %s\n, Total Orders: %.2f", c.Id, c.LastName, c.FirstName, c.PhoneNumber, c.Email, c.Address, c.TotalOrders)
}

func (c Customer) toCSVRecord() []string {
	return []string{
		c.FirstName,
		c.LastName,
		c.PhoneNumber,
		c.Email,
		c.Address,
		fmt.Sprintf("%.2f €", c.TotalOrders),
	}
}

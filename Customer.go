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
	var totalOrders float32
	if c.TotalOrders == 0 {
		totalOrders = 0
	} else {
		totalOrders = c.TotalOrders
	}
	fmt.Printf("Id: %d\nNom: %s\nPrénom: %s\nNuméro de Téléphone: %s\nEmail: %s\nAdresse: %s\nTotal Orders: %.2f€\n\n", c.Id, c.LastName, c.FirstName, c.PhoneNumber, c.Email, c.Address, totalOrders)
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

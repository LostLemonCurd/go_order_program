package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	customers, err := getAllCustomersFromDb(db)
	products, err := getAllProductsFromDb(db)
	orders, err := getAllOrdersFromDb(db)

	// fmt.Println("customers: ", customers)
	// fmt.Println("products: ", products)
	// fmt.Println("orders: ", orders)

	if err != nil {
		panic(err.Error())
	}

	var userChoice int
	for userChoice != 7 {
		userChoice = displayMenu()
		switch userChoice {
		case 1:
			createOneProduct(db)
			products, err = getAllProductsFromDb(db)
			if err != nil {
				panic(err.Error())
			}
		case 2:
			fmt.Println("Voici les produits disponibles: ")
			for _, product := range products {
				product.toString()
			}
		case 3:
			var productId int
			productId = getIntInput("Entrez l'ID du produit à supprimer: ")
			err := modifyOneProductFromDb(db, productId)
			if err != nil {
				panic(err.Error())
			}
			products, err = getAllProductsFromDb(db)
			if err != nil {
				panic(err.Error())
			}
		case 4:
			var productId int
			productId = getIntInput("Entrez l'ID du produit à supprimer: ")
			err := deleteOneProductFromDb(db, productId)
			if err != nil {
				panic(err.Error())
			}
			products, err = getAllProductsFromDb(db)
			if err != nil {
				panic(err.Error())
			}

		case 5:
			writeProductsToCSV(products)
		case 6:
			createOneCustomer(db)
			customers, err = getAllCustomersFromDb(db)
			if err != nil {
				panic(err.Error())
			}
		case 7:
			fmt.Println("Voici les clients: ")
			for _, customer := range customers {
				customer.toString()
			}
		case 8:
			var customerId int
			customerId = getIntInput("Entrez l'ID du client à modifier: ")
			err := modifyOneCustomerFromDb(db, customerId)
			if err != nil {
				panic(err.Error())
			}
			customers, err = getAllCustomersFromDb(db)
			if err != nil {
				panic(err.Error())
			}
		case 9:
			writeCustomersToCSV(customers)
		case 10:
			err := createOneOrder(db, customers, products)
			if err != nil {
				return
			}

			orders, err = getAllOrdersFromDb(db)
			products, err = getAllProductsFromDb(db)
			if err != nil {
				panic(err.Error())
			}

		case 11:
			writeOrdersToCSV(orders)
		case 12:
			exitMenu()
		default:
			exitMenu()
		}
	}
}
func displayMenu() int {
	var userChoice int
	fmt.Printf("\n1- Ajouter un produit\n2- Afficher tous les produits\n3- Modifier un produit\n4- Supprimer un produit \n5- Exporter l'ensemble des produits sous forme csv\n6- Ajouter un client\n7- Afficher tous les clients\n8- Modifier un client\n9- Exporter l'ensemble des clients sous forme csv\n10- Effectuer une commande\n11- Exporter l'ensemble des commandes\n12- Quitter \n: ")
	fmt.Scanln(&userChoice)
	return userChoice
}

func exitMenu() {
	os.Exit(0)
}

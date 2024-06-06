package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func createOneProduct(db *sql.DB) Product {
	var title, description string
	var quantity int
	var price float32

	title = getTextInput("Entrez le titre du produit: ")
	description = getTextInput("Entrez la description du produit: ")
	quantity = getIntInput("Entrez la quantité de produit: ")
	price = getFloatInput("Entrez le prix du produit: ")

	product := Product{Title: title, Description: description, Quantity: quantity, Price: price}
	insertOneProductIntoDb(db, product)
	return product
}

func insertOneProductIntoDb(db *sql.DB, product Product) error {
	// Prepare an SQL statement
	stmt, err := db.Prepare("INSERT INTO products (title, description, quantity, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Execute the statement with the product's data
	_, err = stmt.Exec(product.Title, product.Description, product.Quantity, product.Price)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Produit inséré avec succès")
	return nil
}

func getAllProductsFromDb(db *sql.DB) ([]Product, error) {
	var products []Product
	// Prepare the statement
	stmt, err := db.Prepare(`
		SELECT id, title, description, quantity, price, isActive
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and create Customer objects
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Title, &product.Description, &product.Quantity, &product.Price, &product.IsActive)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func deleteOneProductFromDb(db *sql.DB, productId int) error {
	query := "UPDATE products SET isActive = false WHERE id = ?"
	result, err := db.Exec(query, productId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no product found with id %d", productId)
	}

	fmt.Printf("Marked product with id %d as inactive\n", productId)
	return nil
}

func modifyOneProductFromDb(db *sql.DB, productId int) error {
	var modifiedFields []string
	var modifiedValues []any
	var userChoice int
	for userChoice != 7 {
		userChoice = displayProductMenu()
		switch userChoice {
		case 1:
			title := getTextInput("Entrez le titre du produit: ")
			modifiedFields = append(modifiedFields, "title = ?")
			modifiedValues = append(modifiedValues, title)
		case 2:
			description := getTextInput("Entrez la description du produit: ")
			modifiedFields = append(modifiedFields, "description = ?")
			modifiedValues = append(modifiedValues, description)
		case 3:
			quantity := getStringInput("Entrez la quantité du produit: ")
			modifiedFields = append(modifiedFields, "quantity = ?")
			modifiedValues = append(modifiedValues, quantity)
		case 4:
			price := getStringInput("Entrez le prix du produit: ")
			modifiedFields = append(modifiedFields, "price = ?")
			modifiedValues = append(modifiedValues, price)
		case 5:
			if len(modifiedFields) == 0 {
				fmt.Println("No fields have been modified.")
				continue
			}
			modifiedFieldQuery := strings.Join(modifiedFields, ", ")

			query := fmt.Sprintf("UPDATE products SET %s WHERE id = ?", modifiedFieldQuery)
			modifiedValues = append(modifiedValues, productId)
			result, err := db.Exec(query, modifiedValues...)
			if err != nil {
				return err
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return err
			}

			if rowsAffected == 0 {
				return fmt.Errorf("no product found with id %d", productId)
			}

			fmt.Printf("Updated product with id %d\n", productId)
			return nil
		case 6:
			return nil
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
	return nil
}

func displayProductMenu() int {
	var userChoice int
	fmt.Printf("Que souhaitez-vous modifier? \n1- Le titre\n2- La description\n3- La quantité \n4- Le prix \n5- Valider les modifications\n6- Annuler \n: ")
	fmt.Scanln(&userChoice)
	return userChoice
}

func updateOneProductQuantity(db *sql.DB, order Order) error {
	// Update the quantity of the product in question
	updateStmt, err := db.Prepare("UPDATE products SET quantity = quantity - ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer updateStmt.Close()

	_, err = updateStmt.Exec(order.Quantity, order.Product.Id)
	if err != nil {
		return err
	}

	return nil
}

func writeProductsToCSV(products []Product) {
	csvFile, err := os.Create("products.csv")
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	headers := []string{"Titre", "Description", "Quantité", "Prix", "Supprimé"}
	csvwriter.Write(headers)
	for _, product := range products {
		_ = csvwriter.Write(product.toCSVRecord())
	}
	defer csvwriter.Flush()
	fmt.Print("Le fichier a bien été créé \n")

}

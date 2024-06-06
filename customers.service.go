package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func createOneCustomer(db *sql.DB) Customer {
	var lastName string
	var firstName string
	var phoneNumber string
	var email string
	var address string

	lastName = getTextInput("Entrez le nom du client: ")
	firstName = getTextInput("Entrez le prénom du client: ")
	phoneNumber = getStringInput("Entrez le téléphone du client: ")
	address = getTextInput("Entrez l'adresse du client: ")
	email = getStringInput("Entrez l'email du client: ")

	customer := Customer{LastName: lastName, FirstName: firstName, PhoneNumber: phoneNumber, Address: address, Email: email}
	_, err := insertNewCustomerIntoDb(db, customer)
	if err != nil {
		return Customer{}
	}
	return customer
}

func insertNewCustomerIntoDb(db *sql.DB, customer Customer) (int, error) {
	// Prepare an SQL statement
	stmt, err := db.Prepare("INSERT INTO customers (firstname, lastname, phonenumber, address, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(customer.LastName, customer.FirstName, customer.PhoneNumber, customer.Address, customer.Email)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func getAllCustomersFromDb(db *sql.DB) ([]Customer, error) {
	var customers []Customer
	// Prepare the statement
	stmt, err := db.Prepare(`
		SELECT c.id, c.firstname, c.lastname, c.phonenumber, c.address, c.email,
			   coalesce(total_prices.total_price, 0) as total_price
		FROM customers c
		LEFT JOIN (
			SELECT customer_id, SUM(price) as total_price
			FROM orders
			GROUP BY customer_id
		) total_prices ON c.id = total_prices.customer_id
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
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.LastName, &customer.FirstName, &customer.PhoneNumber, &customer.Address, &customer.Email, &customer.TotalOrders)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func modifyOneCustomerFromDb(db *sql.DB, customerId int) error {
	var modifiedFields []string
	var modifiedValues []any
	var userChoice int
	for userChoice != 7 {
		userChoice = displayCustomerMenu()
		switch userChoice {
		case 1:
			lastName := getTextInput("Entrez le nom de l'employé: ")
			modifiedFields = append(modifiedFields, "lastName = ?")
			modifiedValues = append(modifiedValues, lastName)
		case 2:
			firstName := getTextInput("Entrez le prénom de l'employé: ")
			modifiedFields = append(modifiedFields, "firstName = ?")
			modifiedValues = append(modifiedValues, firstName)
		case 3:
			phoneNumber := getStringInput("Entrez le numéro de téléphone de l'employé: ")
			modifiedFields = append(modifiedFields, "phoneNumber = ?")
			modifiedValues = append(modifiedValues, phoneNumber)
		case 4:
			address := getStringInput("Entrez l'adresse de l'employé: ")
			modifiedFields = append(modifiedFields, "address = ?")
			modifiedValues = append(modifiedValues, address)
		case 5:
			email := getStringInput("Entrez l'email de l'employé: ")
			modifiedFields = append(modifiedFields, "email = ?")
			modifiedValues = append(modifiedValues, email)
		case 6:
			if len(modifiedFields) == 0 {
				fmt.Println("No fields have been modified.")
				continue
			}
			modifiedFieldQuery := strings.Join(modifiedFields, ", ")

			query := fmt.Sprintf("UPDATE customers SET %s WHERE id = ?", modifiedFieldQuery)
			modifiedValues = append(modifiedValues, customerId)
			result, err := db.Exec(query, modifiedValues...)
			if err != nil {
				return err
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return err
			}

			if rowsAffected == 0 {
				return fmt.Errorf("no product found with id %d", customerId)
			}

			fmt.Printf("Updated product with id %d\n", customerId)
			return nil
		case 7:
			return nil
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
	return nil
}

func displayCustomerMenu() int {
	var userChoice int
	fmt.Printf("Que souhaitez-vous modifier? \n1- Le nom de famille\n2- Le prénom\n3- Le numéro de téléphone \n4- L'adresse postale \n5- L'adresse mail \n6- Valider les modifications\n7- Annuler \n: ")
	fmt.Scanln(&userChoice)
	return userChoice
}

func writeCustomersToCSV(customers []Customer) {
	csvFile, err := os.Create("customers.csv")
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	headers := []string{"Nom", "Prénom", "Téléphone", "Email", "Adresse", "Total des Commandes"}
	csvwriter.Write(headers)
	for _, customer := range customers {
		_ = csvwriter.Write(customer.toCSVRecord())
	}
	defer csvwriter.Flush()
	fmt.Print("Le fichier a bien été créé \n")

}

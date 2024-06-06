package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func createOneOrder(db *sql.DB, customers []Customer, products []Product) error {
	var quantity int
	var customerId int
	var productId int
	var orderCustomer Customer
	var orderProduct Product

	quantity = getIntInput("Entrez la quantité: ")
	customerId = getIntInput("Entrez l'id du client: ")
	productId = getIntInput("Entrez l'id du produit: ")

	for _, customer := range customers {
		if customer.Id == customerId {
			orderCustomer = customer
		}
	}
	for _, product := range products {
		if product.Id == productId {
			orderProduct = product
		}
	}

	if quantity > orderProduct.Quantity {
		err := errors.New("Not enough products sorry")
		fmt.Println(err)
		return err
	}
	fmt.Println("orderProduct price", orderProduct.Price)
	orderPrice := float32(quantity) * orderProduct.Price

	order := Order{Quantity: quantity, Customer: orderCustomer, Product: orderProduct, Price: orderPrice}

	fmt.Println(order)
	err := insertOneOrderIntoDb(db, order)
	if err != nil {
		return err
	}
	pdfFilePath := fmt.Sprintf("orders/order_%s.pdf", time.Now().Format("20060102_150405"))
	err = createOrderPDF(order, pdfFilePath)
	if err != nil {
		return err
	}

	return nil
}

func insertOneOrderIntoDb(db *sql.DB, order Order) error {
	stmt, err := db.Prepare("INSERT INTO orders (customer_id, product_id, quantity, price, createdAt) VALUES (?, ?, ?, ?, NOW())")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(order.Customer.Id, order.Product.Id, order.Quantity, order.Price)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	err = updateOneProductQuantity(db, order)
	if err != nil {
		return err
	}

	return nil
}

func getAllOrdersFromDb(db *sql.DB) ([]Order, error) {
	// Prepare the SQL query to fetch orders along with customer and product details
	query := `
        SELECT o.id, o.customer_id, o.product_id, o.quantity, o.price, o.createdAt,
               p.id, p.title, p.description, p.quantity, p.price, p.isActive,
               c.id, c.lastName, c.firstName, c.phoneNumber, c.email, c.address
        FROM orders o
        JOIN products p ON o.product_id = p.id
        JOIN customers c ON o.customer_id = c.id
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the fetched orders
	var orders []Order

	// Iterate over the query results
	for rows.Next() {
		var order Order

		// Scan the order details
		err := rows.Scan(
			&order.Id, &order.Customer.Id, &order.Product.Id, &order.Quantity, &order.Price, &order.CreatedAt,
			&order.Product.Id, &order.Product.Title, &order.Product.Description, &order.Product.Quantity, &order.Product.Price, &order.Product.IsActive,
			&order.Customer.Id, &order.Customer.LastName, &order.Customer.FirstName, &order.Customer.PhoneNumber, &order.Customer.Email, &order.Customer.Address,
		)
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		// Append the order to the slice
		orders = append(orders, order)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func writeOrdersToCSV(orders []Order) {
	csvFile, err := os.Create("orders.csv")
	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	headers := []string{"Id", "Customer first name", "Nom du produit", "Quantité", "Prix", "Date de Création"}
	csvwriter.Write(headers)
	for _, order := range orders {
		_ = csvwriter.Write(order.toCSVRecord())
	}
	defer csvwriter.Flush()
	fmt.Print("Le fichier a bien été créé \n")

}

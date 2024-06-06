package main

import (
	"fmt"
	"github.com/go-pdf/fpdf"
	"os"
	"path/filepath"
)

func createOrderPDF(order Order, filePath string) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Confirmation de commande")
	pdf.Ln(14)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Informations du client")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Client: %s %s", order.Customer.FirstName, order.Customer.LastName))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Telephone: %s", order.Customer.PhoneNumber))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Adresse: %s", order.Customer.Address))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", order.Customer.Email))
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Informations du produit")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Product: %s", order.Product.Title))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Quantite: %d", order.Quantity))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Prix unitaire: %.2f", order.Product.Price))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Prix total: %.2f", order.Price))
	pdf.Ln(10)

	err := pdf.OutputFileAndClose(filePath)
	SummaryCompare(err, filePath)
	if err != nil {
		return err
	}
	return nil
}

func SummaryCompare(err error, fileStr string) {
	if err == nil {
		err = referenceCompare(fileStr)
	}
	if err == nil {
		fileStr = filepath.ToSlash(fileStr)
		fmt.Printf("Successfully generated %s\n", fileStr)
	} else {
		fmt.Println(err)
	}
}

func referenceCompare(fileStr string) (err error) {
	var refFileStr, refDirStr, dirStr, baseFileStr string
	dirStr, baseFileStr = filepath.Split(fileStr)
	refDirStr = filepath.Join(dirStr, "reference")
	err = os.MkdirAll(refDirStr, 0755)
	if err == nil {
		refFileStr = filepath.Join(refDirStr, baseFileStr)
		err = fpdf.ComparePDFFiles(fileStr, refFileStr, false)
	}
	return
}

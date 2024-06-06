package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to get a string input
func getStringInput(prompt string) string {
	var input string
	fmt.Printf(prompt)
	fmt.Scanln(&input)
	return input
}

func getTextInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Function to get an integer input
func getIntInput(prompt string) int {
	for {
		fmt.Printf(prompt)
		var input string
		fmt.Scanln(&input)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}
		fmt.Println("Veuillez entrer un nombre entier valide.")
	}
}

// Function to get a float input
func getFloatInput(prompt string) float32 {
	for {
		fmt.Printf(prompt)
		var input string
		fmt.Scanln(&input)
		value, err := strconv.ParseFloat(input, 32)
		if err == nil {
			return float32(value)
		}
		fmt.Println("Veuillez entrer un nombre décimal valide.")
	}
}

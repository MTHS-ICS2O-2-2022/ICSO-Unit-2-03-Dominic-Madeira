// Created by: Dominic M.
// Created on: March 2023
//
// This program does addition

package main

import "fmt"

func main() {
	// This function does addition
	var houseNumber string
	var streetNumber int

	// input
	fmt.Println("This program finds your address.")
	fmt.Println()
	fmt.Print("Enter your house number: ")
	fmt.Scanln(&houseNumber)
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetNumber)

	// output
	fmt.Println("Your address is:", houseNumber, streetNumber, ".")

	fmt.Println("\nDone.")
}

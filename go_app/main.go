// Created by: Dominic M.
// Created on: March 2023
//
// This program does addition

package main

import "fmt"

func main() {
	// This function does addition
	var houseNumber int
	var streetName string

	// input
	fmt.Println("This program finds your address.")
	fmt.Println()
	fmt.Print("Enter your house number: ")
	fmt.Scanln(&houseNumber)
	fmt.Println()
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetName)

	// output
	fmt.Println()
	fmt.Println("Your address is:", houseNumber, streetName)

	fmt.Println("\nDone.")
}

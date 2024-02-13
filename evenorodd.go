// Write a program to check whether a given number is even or odd.
package main

import "fmt"

func evenOrOdd() {
	var input int
	var result string

	fmt.Print("Please enter a number: ")
	fmt.Scanln(&input)

	if (input % 2 == 0) {
		result = "even"
	} else {
		result = "odd"
	}

	fmt.Printf("%d is %s", input, result)
	fmt.Println()
}
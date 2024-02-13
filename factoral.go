// Write a program to find the factorial of a given number.
package main

import "fmt"

func factorial() {
	var input int
	
	
	fmt.Print("Please enter a number: ")
	fmt.Scanln(&input)
	sum := input

	for i := 1; i < input; i++ {
		sum *= i
	}

	fmt.Printf("Factorial of %d is %d", input, sum)
	fmt.Println()
}
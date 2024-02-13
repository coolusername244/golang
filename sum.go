// Write a program to find the sum of two numbers entered by the user.
package main

import "fmt"

func sum() {
	var num1 int
	var num2 int

	fmt.Print("Num 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Num 2: ")
	fmt.Scanln(&num2)

	var sum int = num1 + num2
	// The type of sum will always be an int so we can infer that type, see below
	// sum := num1 + num2

	fmt.Printf("The sum of %d and %d is %d", num1, num2, sum)
	fmt.Println()
}
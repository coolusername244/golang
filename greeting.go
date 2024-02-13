// Write a program that takes user input (name) and prints a greeting message to that user.
package main

import "fmt"

func greeting() {
	fmt.Println("Hello! What is your name? ")

	var name string

	fmt.Scanln(&name)

	fmt.Println("Hello " + name)
}
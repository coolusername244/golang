// Write a program to check whether a given string is a palindrome or not.
package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome() {
	var input string

	fmt.Print("Enter word to check: ")
	fmt.Scanln(&input)

	var check = Reverse(input)

	if (input == check) {
		fmt.Printf("%s is a palindrome!\n", input)
	} else {
		fmt.Printf("%s is not a palindrome!\n", input)
	}
}
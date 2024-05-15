package main

import (
	"fmt"
)

func main() {
	var n int

	// Prompt the user to enter a positive integer
	fmt.Print("Enter a positive integer: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	// Calculate the sum of even numbers from 1 to n
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	// Print the result
	fmt.Printf("The sum of all even numbers from 1 to %d is: %d\n", n, sum)
}

package main

/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Exercise 002")

	// reading input from console
	var input int
	fmt.Print("Please enter a number : ")
	_, err := fmt.Scanln(&input)

	// checking for error and low 0
	if err != nil {
		log.Fatal("Please enter a number")
	}

	if input < 0 {
		log.Fatal("Please enter a positiv number.")
	}

	// uint64 because it can get big
	var result uint64 = 1

	for i := 1; i <= input; i++ {
		fmt.Print(i)
		result *= uint64(i)
		fmt.Println(result)
	}

	fmt.Printf("Factorial of %d = %d", input, result)
}

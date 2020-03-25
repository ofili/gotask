package main

import (
	"fmt"
)

func convertDecimalToBinary(number int) int {
	binary := 0
	counter := 1
	remainder := 0

	for number != 0 {
		remainder = number % 2
		number = number / 2
		binary += remainder * counter
		counter *= 10

	}
	return binary
}

func main() {
	var decimal int
	fmt.Println("This program let's you convert Decimal numbers to Binary ***")
	fmt.Println("*")
	fmt.Print("Enter Decimal Number: ")
	fmt.Scanln(&decimal)

	fmt.Printf("The Binary is:  %d", convertDecimalToBinary(decimal))

}

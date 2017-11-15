package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter second number: ")
	fmt.Scanf("%d", &b)
	fmt.Printf("Sum of %d and %d is %d", a, b, a+b)
}

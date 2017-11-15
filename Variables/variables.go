/*
Variables and assignments in Go
int int8 int16 int32 int64
float32 float64
string
*/

package main

import "fmt"

func main() {
	x := 'c' // Declaration + initialization - auto typing
	var y int   // Declaration - value will be zero
	y = 10      // Assignment
	fmt.Printf("%T\n", x) // Print type of a variable
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}

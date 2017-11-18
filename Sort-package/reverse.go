package main

import "fmt"
import "sort"

type students []string

func main() {
	s1 := students{"Abc", "Xyz", "Pow", "12qwe"}
	sort.Sort(sort.Reverse(sort.StringSlice(s1)))
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
}

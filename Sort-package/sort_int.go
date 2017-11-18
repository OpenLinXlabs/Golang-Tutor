package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{1, 6, 8, 34, 67, 678, 234, 12, 39}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}

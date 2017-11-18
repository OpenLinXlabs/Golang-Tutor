package main

import (
	"fmt"
	"sort"
)

type students []string

func (s students) Len() int           { return len(s) }
func (s students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s students) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	s1 := students{"Ansil", "Shihab", "Alif", "Shibili"}
	sort.Sort(s1)
	fmt.Println(s1)
}

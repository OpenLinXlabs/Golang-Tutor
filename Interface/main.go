package main

import "fmt"

type Square struct {
	Side float64
}
type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z.area())
}
func main() {
	s1 := Square{10}
	info(s1)
	c1 := Circle{3}
	info(c1)
}

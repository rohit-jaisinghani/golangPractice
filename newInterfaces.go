package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type rectangle struct {
	length, breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
func print(s shape) {
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s.area())
	fmt.Printf("%v\n", s.perimeter())
}

type empty interface {
}
type person struct {
	info interface{}
}

func main() {
	var emptyintf interface{}
	emptyintf = 5
	c:=circle{radius:4}
	fmt.Println(emptyintf)
	fmt.Println(c.area())
	var s shape = circle{radius: 5.0}
	fmt.Printf("Circle Area:%v\n", s.area())
	s = rectangle{length: 5.0, breadth: 5.0}
	fmt.Printf("Rectangle Area:%v\n", s.area())

}

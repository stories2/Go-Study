package main

import (
	"fmt"
	"math"
)

type Circle struct {
	// x float64
	// y float64
	// z float64
	x, y, z float64
}

type Person struct {
	name string
}

func (p *Person) SayMyName() {
	fmt.Println("My name is ", p.name)
}

type Android struct {
	Person
	model string
}

func main() {
	// c := new (Circle)
	c := Circle{x: 1, y: 2, z: 3}

	fmt.Println(c.x, c.y, c.z)

	fmt.Println(circleArea(c))

	fmt.Println(circleAreaRef(&c), c.z)

	fmt.Println(c.area())

	a := new(Android)
	// fmt.Println()
	a.Person.name = "hell"
	a.Person.SayMyName()
}

func circleArea(c Circle) float64 {
	return math.Pi * c.y * c.y
}

func circleAreaRef(c *Circle) float64 {
	c.z = math.Pi * c.y * c.y
	return c.z
}

func (c *Circle) area() float64 {
	return math.Pi * c.y * c.y
}

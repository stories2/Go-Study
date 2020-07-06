package main

import "fmt"

var c string = "Hell world"

func main() {
	var x string = "Hell World"
	fmt.Println(x)
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = x + "thrid"
	fmt.Println(x)

	x = "hell"
	var y string = "world"
	fmt.Println(x == y)
	y = "hell"
	fmt.Println(x == y)

	a := "Hell world"
	fmt.Println(a)

	var b = "Hell world"
	fmt.Println(b)

	fmt.Println(c)

	const e = "Hell world"
	// e = "test"

	var (
		f = 10
		g = 20
		h = 30
	)
	fmt.Println(f, g, h)

	fmt.Println("Input: ")
	var input float64
	fmt.Scanf("%f\n", &input)
	output := input * 2
	fmt.Println("*2 = ", output)
}

func showX() {
	fmt.Println(c)
	// fmt.Println(x)
}

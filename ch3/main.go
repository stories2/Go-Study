package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1.0+1.0)

	fmt.Println(len("Hell World"))
	fmt.Println("Hell World"[1])
	fmt.Println("Hell" + " World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

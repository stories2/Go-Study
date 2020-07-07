package main

import "fmt"

func main() {
	x := 5
	xPointer(&x)
	fmt.Println(x)
}

func xPointer(xPtr *int) {
	*xPtr = 0
}

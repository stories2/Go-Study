package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i = 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("짝")
		} else {
			fmt.Println("홀")
		}
	}

	for i = 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("-0")
			break
		case 1:
			fmt.Println("-1")
			break
		default:
			fmt.Println("-*")
		}
	}
}

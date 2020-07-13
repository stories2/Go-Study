package main

import "fmt"
import "./math"

func main() {
	xs := []float64{10, 20, 30}

	avg := math.Avg(xs)

	fmt.Println("avg", avg)
}
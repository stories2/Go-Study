package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 10
	y[1] = 20
	y[2] = 30
	y[3] = 40
	y[4] = 50

	var sum float64 = 0

	for i := 0; i < len(y); i++ {
		sum += y[i]
	}
	fmt.Println(sum / float64(len(y)))

	sum = 0
	for _, value := range y {
		sum += value
	}
	fmt.Println(sum / float64(len(y)))

	z := [5]float64{50, 60, 70, 80, 90}

	fmt.Println(z)

	v := make([]float64, 5)
	fmt.Println(v)
	v = make([]float64, 5, 10)
	fmt.Println(v)

	arr := []float64{1, 2, 3, 4, 5}
	t := arr[1:4]
	fmt.Println(t)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2)

	copy1 := []int{1, 2, 3}
	copy2 := make([]int, 2)
	copy(copy2, copy1)
	fmt.Println(copy2)

	keyVal := make(map[string]int)
	keyVal["hell"] = 10
	fmt.Println(keyVal)

	if name, ok := keyVal["world"]; !ok {
		fmt.Println(name, ok)
	}

	el := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}

	fmt.Println(el)

	if item, ok := el["H"]; ok {
		fmt.Println(item["name"], item["state"])
	}
}

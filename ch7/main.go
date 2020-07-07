package main

import "fmt"

func main() {
	a := []float64{10, 20, 30, 40, 50}

	fmt.Println(avg(a))

	b, c := multiReturn()
	fmt.Println(b, c)

	fmt.Println(sum(1, 2, 3, 4, 5))

	add := func(a int, b int) int {
		return a + b
	}

	fmt.Println(add(4, 5))

	x := 2
	des := func() int {
		x--
		return x
	}

	fmt.Println(des(), des())

	even := makeEvenGen()
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())

	fmt.Println(fact(5))

	deferTest()

	defer func() {
		str := recover()
		fmt.Println("recover", str)
	}()

	panic("PAN")
}

func deferTest() {
	defer third()
	first()
	second()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func third() {
	fmt.Println("thrid")
}

func fact(x int) int {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}

func makeEvenGen() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func sum(a ...int) int {
	sum := 0
	for _, val := range a {
		sum += val
	}
	return sum
}

func avg(a []float64) float64 {
	sum := 0.0

	for _, val := range a {
		sum += val
	}

	return sum / float64(len(a))
}

func multiReturn() (int, int) {
	return 5, 6
}

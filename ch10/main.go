package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	go foo(i)
	// }
	// var input string
	// fmt.Scanln(&input)

	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "C1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "C2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("msg1", msg1)
			case msg2 := <-c2:
				fmt.Println("msg2", msg2)
			case <-time.After(time.Second):
				fmt.Println("msg time sec")
				// default:
				// 	fmt.Println("default5")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func foo(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(n, ":", i)

		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

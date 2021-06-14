package main

import "fmt"

func runSimpleChannel() {
	c := make(chan int64)
	defer close(c)

	go func(c chan int64) {
		for i := 0; i < 10; i++ {
			c <- int64(i)
		}
	}(c)

	fmt.Println("main...")
	for i := 0; i < 10; i++ {
		v := <- c
		fmt.Println(v)
	}
}

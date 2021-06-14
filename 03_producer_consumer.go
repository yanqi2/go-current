package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runProducerConsumer() {
	c := make(chan int64, 1000)
	defer close(c)

	go runProducer(c)
	go runConsumer(1, c)
	var x string
	fmt.Scan(&x)
}

func runProducer(c chan<- int64) {
	for {
		i := rand.Intn(3) + 1
		time.Sleep(time.Duration(i) * time.Second)
		t := time.Now().Unix()
		c <- t
		fmt.Println("produce: ", t)
	}
}

func runConsumer(n int, c <-chan int64) {
	for {
		i := rand.Intn(3) + 6
		time.Sleep(time.Duration(i) * time.Second)
		t := <-c
		fmt.Println("consume_", n, ": ", t)
	}
}

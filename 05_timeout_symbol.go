package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runTimeoutSymbol() {
	workChan := make(chan int)
	go work(workChan)
	select {
	case <-workChan:
		fmt.Println("done")
		break
	case <- time.After(1*time.Second):
		fmt.Println("time out...")
		break
	}
}

func work(c chan<- int) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("work")
	c <- 1
}


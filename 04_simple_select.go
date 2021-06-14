package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runSimpleSelect() {
	aChan := make(chan int)
	bChan := make(chan int)
	cChan := make(chan int)
	go downloadFromA(aChan)
	go downloadFromB(bChan)
	go downloadFromC(cChan)
	select {
	case <- aChan:
		fmt.Println("a done")
		break
	case <- bChan:
		fmt.Println("b done")
		break
	case <- cChan:
		fmt.Println("c done")
		break
	}
}

func downloadFromA(c chan<- int) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("a work")
	c <- 1
}

func downloadFromB(c chan<- int) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("b work")
	c <- 1
}

func downloadFromC(c chan<- int) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("c work")
	c <- 1
}

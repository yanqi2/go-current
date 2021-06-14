package main

import (
	"fmt"
	"time"
)

/*
创建一个简单的协程执行 doSomething 方法
*/
func runSimpleGoroutine() {
	go doSomething()
	fmt.Println("main...")
	time.Sleep(1 * time.Second)
}

func doSomething(){
	fmt.Println("do something...")
}

func runSimpleGoroutineV2() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i:=0;i<10;i++{
			c <- i
			time.Sleep(1 * time.Second)
		}
	}()
	for i := range c {
		fmt.Println(i)
	}
}
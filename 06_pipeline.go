package main

import (
	"fmt"
	"time"
)

func runPipeline() {
	c1 := process01(10)
	c2 := process02(c1)
	c3 := process03(c2)

	for c := range c3 {
		fmt.Println(c)
	}
}

func process01(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- fmt.Sprint("配件", i)
			time.Sleep(1 * time.Second)
		}
	}()
	return out
}

func process02(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装（" + c + "）"
		}
	}()
	return out
}

func process03(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包（" + c + "）"
		}
	}()
	return out
}

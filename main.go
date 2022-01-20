package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
}

func main() {
	c := make(chan int)
	go countToTen(c)
	// a := <-c			 // <-c == for{} == Program is waiting
	fmt.Println("blocking...")
	var a []int
	for {
		a = append(a, <-c)
	}
	fmt.Printf("received %d\n", a)
	fmt.Println(a)
}

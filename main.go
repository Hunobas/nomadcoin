package main

import (
	"fmt"
	"time"
)

func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
	close(c)
}

func receive(c <-chan int) []int {
	var a []int
	for {
		ac, ok := <-c
		if !ok {
			fmt.Println("we are done.")
			break
		}
		a = append(a, ac)
	}
	return a
}

func main() {
	c := make(chan int)
	go countToTen(c)
	// a := <-c			 // <-c == for{} == Program is waiting
	fmt.Println("blocking...")
	a := receive(c)
	fmt.Printf("received %d\n", a)
	fmt.Println(a)
}

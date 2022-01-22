package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf(">>sending %d<<\n", i)
		c <- i // blocking operation
		fmt.Printf(">>sent %d<<\n", i)
	}
	close(c)
}

func receive(c <-chan int) []int {
	var a []int
	for {
		time.Sleep(2 * time.Second)
		ac, ok := <-c // blocking operation
		if !ok {
			fmt.Println("we are done.")
			break
		}
		fmt.Printf("|| received %d ||\n", ac)
		a = append(a, ac)
	}
	return a
}

func main() {
	c := make(chan int, 10)
	go send(c)
	// a := <-c			 // <-c == for{} == Program is waiting
	fmt.Println("blocking...")
	a := receive(c)
	fmt.Printf("received %d\n", a)
	fmt.Println(a)
}

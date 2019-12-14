package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		fmt.Println("func : ", v)
		time.Sleep(100 * time.Millisecond)
	}
	c <- sum // send sum to c
	//close(c) not working olny one  send .
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 1}
	//fmt.Println(s[:len(s)/2])
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	fmt.Println("a")
	go sum(s[len(s)/2:], c)
	fmt.Println("b")
	x, y := <-c, <-c // receive from c
	fmt.Println("c")
	fmt.Println(x, y, x+y)
}

package main

import (
	"fmt"
	"time"
)

func say(s string, t float32, c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(t) * time.Millisecond)
		fmt.Println(i, s)
	}
	close(c)
}

func main() {
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	go say("world", 2, a)
	go say("hello", 10, c)
	say("Chang", 1, b)

	for {
		_, oka := <-a
		_, okb := <-b
		_, okc := <-c
		if !oka && !okb && !okc {
			fmt.Println("<-- loop broke!")
			break // exit break loop
		}
	}
	fmt.Println("main() stopped")
}

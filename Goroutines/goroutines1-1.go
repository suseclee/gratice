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
}

func main() {
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	go say("world", 2, a)
	go say("hello", 10, c)
	say("Chang", 1, b)
	_, _, _ = <-a, <-b, <-c
	close(a)
	close(b)
	close(c)
}

// fatal error: all goroutines are asleep — deadlock!.
// Seems like all goroutines are asleep or simply no other goroutines are available to schedule.

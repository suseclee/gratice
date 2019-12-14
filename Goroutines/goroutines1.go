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
	c <- "done"
}

func main() {
	ch := make(chan string, 3)
	go say("world", 2, ch)
	go say("hello", 10, ch)
	say("Chang", 1, ch)

	for {
		done := true
		// TODO
		// Get   v, ok  := range ch
		for v := range ch {
			fmt.Println("read value", v, "from ch")
			time.Sleep(2 * time.Second)
			_, ok := <-v
			done = done && ok
		}

		if done {
			break
		}

	}
	close(ch)

}

// fatal error: all goroutines are asleep — deadlock!.
// Seems like all goroutines are asleep or simply no other goroutines are available to schedule.

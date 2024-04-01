package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 1000)
	}
}
func main() {
	var c chan string = make(chan string)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from c1"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	go func() {
		for {
			c2 <- "from c2"
			time.Sleep(time.Millisecond * 3000)
		}
	}()


	go func ()  {
		
	}
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

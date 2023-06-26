package main

import (
	"fmt"
	"time"
)

func ping(char chan string) {
	for i := 0; ; i++ {
		char <- "ping"
	}
}
func pong(char chan string) {
	for i := 0; ; i++ {
		char <- "pong"
	}
}

func print(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go print(c)
	var entrada string
	fmt.Scanln(&entrada)
}

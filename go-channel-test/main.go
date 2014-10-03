package main

import (
	"fmt"
	"time"
)

func ping(c chan<- string) {
	for i := 0; ; i++ {
		message := "ping"
		fmt.Println(message)
		c <- message
	}
}

func receiver(c <-chan string) {
	for {
		<-c
		response := "pong"
		fmt.Println(response)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)
	go ping(c)
	go receiver(c)
	fmt.Scanln()
	print("Exit")
}

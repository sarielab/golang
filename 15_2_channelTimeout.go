package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(i) * time.Second)
	}
}

func receiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println(`received data "`, data, `"\n`)
		case <-time.After(time.Second * 5):
			fmt.Println("timeout no activities after 5 seconds")
			break loop
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().Unix())
	var msgs = make(chan int)
	go sendData(msgs)
	receiveData(msgs)
}

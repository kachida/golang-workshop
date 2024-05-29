package main

import (
	"fmt"
	"time"
)

//A goroutine is a lightweight thread of execution.
func f(from string) {
	for i := 0; i<3 ; i++ {
		time.Sleep(time.Second)
		fmt.Println(from , ":", i)
	}
}

func main() {
	
	//Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
	f("direct")

	//To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f("goroutine 1")
	
	//You can also start a goroutine for an anonymous function call.
	go f("goroutine 2")
	
	//Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
	go f("goroutine 3")

	//When we run this program, we see the output of the blocking call first, then the output of the two goroutines. 
	//The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	time.Sleep(time.Second*20)
	fmt.Println("done")
	//Next we’ll look at a complement to goroutines in concurrent Go programs: channels.
}
package main

import (
	"fmt"
)

func main() {

	n:=10
	c:= make(chan int)  //create a new channel which will be receiving int values
	done := make(chan bool) //create another channel which will receive only a bool value

	go func() {
		for i:=0; i < 1000 ; i++ {
			c <- i //send the int values to channel c
		}
	}()

	for i:= 0; i < n ; i++ {
		go func(x int) {
			for q:= range c {
				fmt.Println("I",x,"q", q)
			}
			done <- true //send the boolean values to done channel
		}(i)
	}

	for i:= 0; i < n ; i++ {
		<-done //here we are sending the value from done to unknown target which will clear the done queue
	}
}
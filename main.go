package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(5 * time.Second) //simulate a slow task
	fmt.Println("Slow Hello", phrase)

	doneChan <- true
}

func main() {
	// go greet("Greet 1")
	// go greet("Greet 2")
	done := make(chan bool)
	go slowGreet("Greet 3", done)
	// go greet("Greet 4")

	<-done
}

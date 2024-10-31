package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) //simulate a slow task
	fmt.Println("Slow Hello", phrase)
}

func main() {
	greet("Greet 1")
	greet("Gree 2")
	slowGreet("Greet 3")
	greet("Greet 4")
}

package main

import (
	"fmt"
	// "rsc.io/quote"
)

// func main() {
// 	fmt.Println(quote.Hello())
// }

// func main() {
// 	fmt.Println("Hello, world")
// }

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}

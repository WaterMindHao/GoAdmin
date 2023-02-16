package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}

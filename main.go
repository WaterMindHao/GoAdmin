package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

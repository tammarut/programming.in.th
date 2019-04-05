package main

import "fmt"

func main() {
	var A int
	fmt.Print("Input A: ")
	fmt.Scanln(&A)

	var B int
	fmt.Print("Input B: ")
	fmt.Scanln(&B)

	fmt.Printf("%d + %d = %d", A, B, A+B)
}

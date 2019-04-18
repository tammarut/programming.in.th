package main

import "fmt"

func main() {
	var box []int

	for i := 0; i < 5; i++ {
		var input int
		fmt.Print("Enter 5 numbers: ")
		fmt.Scanln(&input)
		box = append(box, input)
	}

	fmt.Println(box)
}

package main

import (
	"fmt"
	"log"
)

func main() {
	var box []int

	fmt.Println("Enter 5 numbers")
	for i := 1; i <= 5; i++ {
		var input int
		fmt.Printf("\t %d)", i)
		fmt.Scanln(&input)
		if input < 1 || input > 1000 {
			log.Println("Error: Please key 1-1000!")
		}
		box = append(box, input)
	}

	fmt.Println(box)
}

package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var box []int

	fmt.Println("Enter 5 numbers")
	for i := 1; i <= 5; i++ {
		var input int
		fmt.Printf("\t %d)", i)

		for {
			fmt.Scanln(&input)
			if input >= 1 && input <= 1000 {
				break
			}
			log.Println("Please key 1-1000!")
		}
		box = append(box, input)
	}

	sort.Ints(box)
	fmt.Printf("Min: %d\n", box[0])
	fmt.Printf("Max: %d\n", box[len(box)-1])
}

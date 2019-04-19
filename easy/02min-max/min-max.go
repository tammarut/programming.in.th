package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var size int

	fmt.Print("Enter size of box: ")
	fmt.Scanln(&size)

	box := make([]int, size)

	fmt.Printf("Enter %d numbers\n", size)
	for i := 0; i < len(box); i++ {
		fmt.Printf("\t %d)", i+1)

		for {
			fmt.Scanln(&box[i])
			if box[i] >= 1 && box[i] <= 1000 {
				break
			}
			log.Println("Please key 1-1000!")
		}
	}

	sort.Ints(box)
	fmt.Printf("Min: %d\n", box[0])
	fmt.Printf("Max: %d\n", box[len(box)-1])
}

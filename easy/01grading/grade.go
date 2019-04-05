package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("คะแนนเก็บ: ")
	fmt.Scanln(&a)
	if a >= 0 && a <= 30 {
		fmt.Print("คะแนนกลางภาค: ")
		fmt.Scanln(&b)
		if b >= 0 && b <= 30 {
			fmt.Print(b)
		} else {
			fmt.Println("Please key 0-30!")
		}
	} else {
		fmt.Println("Please key 0-30!")
	}

}

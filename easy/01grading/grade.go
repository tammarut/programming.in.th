package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("คะแนนเก็บ(0-30): ")
	fmt.Scanln(&a)
	if a >= 0 && a <= 30 {
		fmt.Print("คะแนนกลางภาค(0-30): ")
		fmt.Scanln(&b)
		if b >= 0 && b <= 30 {
			fmt.Print("คะแนนปลายภาค(0-40): ")
			fmt.Scanln(&c)
			if (c < 0) || (c > 40) {
				fmt.Println("Please key 0-40")
			}
		} else {
			fmt.Println("Please key 0-30!")
		}
	} else {
		fmt.Println("Please key 0-30!")
	}

	sum := a + b + c
	if sum >= 80 && sum <= 100 {
		fmt.Println("A")
	}
}

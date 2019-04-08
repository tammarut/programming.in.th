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

	fmt.Println(getGrade(a + b + c))

}

func getGrade(sum int) string {
	if sum >= 80 && sum <= 100 {
		return "A"
	}
	if sum >= 75 && sum <= 79 {
		return "B+"
	}
	if sum >= 70 && sum <= 74 {
		return "B"
	}
	if sum >= 65 && sum <= 69 {
		return "C+"
	}
	if sum >= 60 && sum <= 64 {
		return "C"
	}
	if sum >= 55 && sum <= 59 {
		return "D+"
	}
	if sum >= 50 && sum <= 54 {
		return "D"
	}
	return "F"
}

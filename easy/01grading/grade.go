package main

import "fmt"

func main() {
	var a int

	fmt.Print("Key คะแนนเก็บ: ")
	fmt.Scanln(&a)
	if a >= 0 && a <= 30 {

	} else {
		fmt.Println("Please key 0-30!")
	}

}

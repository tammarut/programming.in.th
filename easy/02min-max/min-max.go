package main

//! หาMinimum-Maximum
import (
	"fmt"
	"log"
	"sort"
)

func main() {
	//=> 1) Enter size of box
	var size int
	fmt.Print("Enter size of box: ")
	fmt.Scanln(&size)

	//=> 2) Create slice depend on size
	box := make([]int, size)

	//=> 3) Enter numbers
	fmt.Printf("Enter %d numbers\n", size)
	for i := 0; i < len(box); i++ {
		fmt.Printf("\t %d)", i+1)

		//! handle: ถ้ามากกว่า1 หรือน้อยกว่า1000 ไปต่อ
		for {
			fmt.Scanln(&box[i])
			if box[i] >= 1 && box[i] <= 1000 {
				break
			}
			log.Println("Please key 1-1000!")
		}
	}

	//=> 4) Sort numbers| Ascending
	sort.Ints(box)

	//=> 5) output Min/Max
	fmt.Printf("Min: %d\n", box[0])
	fmt.Printf("Max: %d\n", box[len(box)-1])
}

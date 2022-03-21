package main

import (
	"fmt"
)

func main() {
	
	var numbers [11]int

	for i := range numbers {
		if i % 2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}
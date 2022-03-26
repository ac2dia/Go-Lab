package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	
	// fmt.Println(bytes) // print to bytes
	fmt.Println(string(bytes)) // print to string
}
package main

import (
	"fmt"
)

func main() {
	// colors := make(map[int]string) // create a new map with no values inside

	colors := map[string]string{
		"red": "#ff00000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yelolow"] = "laskasdasd"
	delete(colors, "yellow")
	// delete(colors, 10) // delete key, value in map

	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Printf("Hex code for %s is %s\n", color, hex)
	}
}

/*
# What is Maps?
Maps
- key -> value
- all the same type!!

# Declaring Maps
map[key_type]value_type{

}

- var map_name map[key_type]value_type
- make(map[key_type]value_type)

# Differences between Maps and Structs
- key, value 쌍들의 타입이 같아야함 / value 타입 달라도 상관 없음
- 인덱스 접근 가능 / 인덱스 접근 불가
- 참조 타입 / 값 타입

*/
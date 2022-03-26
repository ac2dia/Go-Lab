package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://pkg.go.dev/net/http")
	if err != nil {
		log.Fatal(err)
		return
	}
	
	fmt.Println(resp)

}
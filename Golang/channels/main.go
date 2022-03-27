package main

import (
	"fmt"
	"net/http"
)

func main() {
	

	links := []string {
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
		"https://www.inflearn.com/",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		return
	}

	fmt.Println(link + " is up!")
}

/*
# Go Routines

go func()
- create a new thread go routine
- And run this function with it

Go Scheduler
- CPU Core
- Go Routine




*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	

	links := []string {
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
		"https://www.inflearn.com/",
		"https://www.udemy.com/",
	}

	c := make(chan string)	
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link) // occurs block!
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- link
		return
	}

	fmt.Println(link + " is up!")
	c <- link
}

/*
# Go Routines

go func()
- create a new thread go routine
- And run this function with it

Go Scheduler
- Go Routines 를 모니터링하는 역할
- One CPU Core 와 Multiple Go Routine 구성

- Multiple CPU Core 와 Multiple Go Rountines


# Channel
channel 을 생성할 때 타입을 지정할 수 있으며 해당 채널을 통해 Go Routine 과 통신하는 데이터도 같은 타입이어야한다.

Sending Data with Channels
- channel <- 5 // 채널로 메시지를 전송
- myNumber <- channel // 채널로 값이 들어올때까지 대기하다가 들어오는 경우 해당 변수에 할당

*/
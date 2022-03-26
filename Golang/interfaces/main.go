package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {

}

type koreanBot struct {

}

func main() {
	eb := englishBot{}
	kb := koreanBot{}

	eb.getGreeting()
	kb.getGreeting()
	
	printGreeting(eb)
	printGreeting(kb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (koreanBot) getGreeting() string {
	return "안녕하세요"
}

/*
# Interfaces

## Problems Without Interfaces
- Interface 가 없다면 메서드 이름이 같을시 파라미터가 다르더라도 에러 발생!
- 클래스 개념이 없기 때문에 메서드 override 불가!



*/
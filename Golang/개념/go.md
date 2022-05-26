# Go 주요 키워드

## defer

- 특정 문장 혹은 함수르 나중에 실행하게 한다.
- 일반적으로 C#, Java 에서 finally 블록처럼 마지막에 clean-up 작업을 위해 사용된다.

```go
package main

import "os"

func main() {
    f, err := os.Open("1.txt")
    if err != nil {
        panic(err)
    }

    // main 마지막에 파일 close 실행
    defer f.Close()

    // 파일 읽기
    bytes := make([]byte, 1024)
    f.Read(bytes)
    println(len(bytes))
}
```

## panic

- 내장함수인 panic() 함수는 현재 함수를 즉시 멈추고, 현재 함수에 있는 defer 키워드를 모두 실행한 후 즉시 리턴한다.
- 이러한 실행 모드 방식은 상위함수에도 똑같이 적용되고, 계속 콜스택을 타고 올라가며 적용된다.
- 마지막에는 프로그램이 에러를 내고 종료하게 된다.

```go
package main

import "os"

func main() {
    // 잘못된 파일명을 넣음
    openFile("Invalid.txt")

    // openFile() 안에서 panic이 실행되면
    // 아래 println 문장은 실행 안됨
    println("Done")
}

func openFile(fn string) {
    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }

    defer f.Close()
}

```

## recover

- 내장함수인 recover() 함수는 panic() 함수에 의한 패닉 상태를 다시 정상 상태로 되돌리는 함수이다.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 잘못된 파일명을 넣음
    openFile("Invalid.txt")

    // recover에 의해
    // 이 문장 실행됨
    println("Done")
}

func openFile(fn string) {
    // defer 함수. panic 호출시 실행됨
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("OPEN ERROR", r)
        }
    }()

    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }

    defer f.Close()
}
```

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

## recover

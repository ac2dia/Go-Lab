package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

/* Questions!
1. How do we run the code in our project?
- go help
- go run main.go
- build, run, fmt, install, get, test

2. What does 'package main' mean?
- package == project == workspace
- Types of Packages
  - Executable (main)
  - Reusable (any other package)
- package 이름이 main 인 경우에만 실행 파일 생성

3. What does 'import "fmt"' mean?
- library
- 다양한 표준 라이브러리: golang.org/pkg

4. What's that 'func' thing?
- function(Method) of General Programming language

5. How is the main.go file organized?
- package declaration
  - package main
- Import other packages
  - import "fmt"
- Declare functions
  - func main() {
	  fmt.Println("hi there")
  }

*/

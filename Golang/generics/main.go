package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz", "hello"}
	fmt.Println(index(ss, "hello"))
}


/*
# Generics
## Type parameters

- func Index[T comparable](s []T, x T) int


## Generic types

``` go
package main

import "fmt"

// values of any type.
type List[T any] struct {
	val  T
}

func main() {
	s := List[string]{"simple"} // [type] <- 타입 초기화
	fmt.Println(s)
}
```


*/
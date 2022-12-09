package main

import (
	"fmt"

	"github.com/hadihammurabi/go-result/result"
)

func IsOdd(num int) (result.Option[bool], result.Option[error]) {
	if num%2 == 0 {
		return result.None[bool](), result.Some(fmt.Errorf("%d is not odd", num))
	}

	return result.Some(true), result.None[error]()
}

func main() {
	res, err := IsOdd(2)
	if err.Ok() {
		panic(err.Get())
	}

	fmt.Println(res.Get())
}

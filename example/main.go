package main

import (
	"fmt"

	"github.com/hadihammurabi/go-result/result"
)

func main() {
	BasicSome()
	BasicNone()
	EqualsSome()
	EqualsNone()

	GlobalUtils()
}

func BasicSome() {
	fmt.Println("BasicSome")
	res := result.Some(10)
	fmt.Println(res.Ok())
	fmt.Println(res.OkOr(0))
	fmt.Println(res)

	fmt.Println()
}

func BasicNone() {
	fmt.Println("BasicNone")
	res := result.None[int]()
	fmt.Println(res.Ok())
	fmt.Println(res.OkOr(0))
	fmt.Println(res)

	fmt.Println()
}

func EqualsSome() {
	fmt.Println("EqualsSome")
	res1 := result.Some(10)
	res2 := result.Some(10)
	fmt.Println(res2.Equals(res1))

	fmt.Println()
}

func EqualsNone() {
	fmt.Println("EqualsNone")
	res1 := result.None[int]()
	res2 := result.Some(1)
	fmt.Println(res2.Equals(res1))

	fmt.Println()
}

func GlobalUtils() {
	fmt.Println("GlobalUtils")
	fmt.Println(result.IsFalsy(false))
	fmt.Println(result.IsTruthy("true"))
	fmt.Println(result.ValueOr("", "default"))

	fmt.Println()
}

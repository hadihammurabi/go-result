# go-result
Result utilities for Go

# Example
```go
import "github.com/hadihammurabi/go-result/result"

func Division(a, b int) result.Option[int] {
	if b == 0 {
		return result.None[int]()
	}

	return result.Some(a / b)
}

func TryDivision(a, b int) {
	try := Division(a, b)
	if !try.Ok() {
		fmt.Printf("%d / %d = failed\n", a, b)
	} else {
		fmt.Printf("%d / %d = %d\n", a, b, try.Get())
	}
}

func main() {
	TryDivision(2, 0) // 2 / 0 = failed
	TryDivision(8, 2) // 8 / 2 = 4
}

```

> see full examples in [example](./example) directory

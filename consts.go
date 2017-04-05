package main

import "fmt"
import "math"

const c string = "const"

func main() {
	fmt.Println(c)

	const n = 50000
	const d = 3e20 / n

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

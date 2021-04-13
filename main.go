package main

import (
	"fmt"
	"lodash-go/lodash"
)

func main() {
	var v float64
	err := lodash.Chain([]float32{2.34, 1.324}).Avg().Value(&v)
	fmt.Println(err, v)
}

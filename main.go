package main

import (
	"fmt"
	"lodash-go/lodash"
)

func main() {
	newStr := ``
	lodash.Chain([]int{1,2,3}).Join(``).ConcatStr(`4`,`5`).Value(&newStr)
	fmt.Println(newStr)
}

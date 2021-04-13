package main

import (
	"fmt"
	"reflect"
)

func main() {
	var int1 int
	int1 = 1
	fmt.Println(reflect.ValueOf(int1).Int())
}

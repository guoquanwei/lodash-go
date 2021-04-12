package lodash

import (
	"errors"
	"fmt"
	"reflect"
)

// TODO
// search split replace trim indexOfStr includesStr substring charAt() toLowerCase() toUpperCase()

func ConcatStr (input interface{}, inputs ...string) (string, error) {
	inputRv := reflect.ValueOf(input)
	if inputRv.Kind().String() != `string` {
		return ``, errors.New(fmt.Sprintf(`%s args.input must be string`, `ConcatStr`))
	}
	inputStr := inputRv.String()
	for _, str :=  range inputs {
		inputStr += str
	}
	return inputStr, nil
}

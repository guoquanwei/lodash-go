package lodash

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotFound       = errors.New("record not found")
	ErrNotImplemented = errors.New("not implemented")
	ErrInvalidData    = errors.New("invalid data")
	ErrInvalidField   = errors.New("invalid field")
	ErrInvalidValue   = errors.New("invalid value")
)

//Reflect Types
//Bool:          "bool",
//Int:           "int",
//Int8:          "int8",
//Int16:         "int16",
//Int32:         "int32",
//Int64:         "int64",
//Uint:          "uint",
//Uint8:         "uint8",
//Uint16:        "uint16",
//Uint32:        "uint32",
//Uint64:        "uint64",
//Uintptr:       "uintptr",
//Float32:       "float32",
//Float64:       "float64",
//Complex64:     "complex64",
//Complex128:    "complex128",
//Array:         "array",
//Chan:          "chan",
//Func:          "func",
//Interface:     "interface",
//Map:           "map",
//Ptr:           "ptr",
//Slice:         "slice",
//String:        "string",
//Struct:        "struct",
//UnsafePointer: "unsafe.Pointer",
var ReflectIntTypes = []string{`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`}
var ReflectFloatTypes = []string{`float32`, `float64`, `complex64`, `complex128`}
var ReflectArrayTypes = []string{`array`, `slice`, `map`}

func chainArgConvert(arg interface{}) (interface{}, bool) {
	useArg := arg
	isChain := false
	switch arg.(type) {
	case *lodash:
		useArg = arg.(*lodash).input
		isChain = true
	}
	return useArg, isChain
}

// output: 外部变量输出地址, input: 入参【真入参/lodash对象】, isChain: 是否是链， result：运算结果
func chainOutputConvert(output interface{}, l interface{}, isChain bool, result interface{}) error {
	if isChain {
		l.(*lodash).input = result
	}
	if !isChain {
		inputKind := reflect.ValueOf(result).Kind().String()
		switch inputKind {
		case `array`, `chan`, `map`, `slice`:
			outputJson, err := json.Marshal(result)
			if err != nil {
				return errors.New(fmt.Sprintf(`output format error: %s`, err.Error()))
			}
			err = json.Unmarshal(outputJson, output)
			if err != nil {
				return errors.New(fmt.Sprintf(`output format error: %s`, err.Error()))
			}
		default:
			reflect.ValueOf(output).Elem().Set(reflect.ValueOf(result))
		}
	}
	return nil
}

func chainOutputNoSlice(output interface{}, l interface{}, isChain bool, result interface{}) {
	if isChain {
		l.(*lodash).input = result
	} else {
		reflect.ValueOf(output).Elem().Set(reflect.ValueOf(result))
	}
}

func CheckKindErr(funcName string, isChain bool, outputKind string, inputKind string) error {
	if !Includes(ReflectArrayTypes, inputKind) {
		return errors.New(fmt.Sprintf(`%s args.input must be slice`, funcName))
	}
	if !isChain {
		if outputKind != `ptr` {
			return errors.New(fmt.Sprintf(`%s args.output must be pointer`, funcName))
		}
	}
	return nil
}

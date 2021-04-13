package lodash

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
)

// TODO
//  random

func Min(input interface{}) (num float64, err error) {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	if !Includes(ReflectArrayTypes, inputRv.Kind().String()) {
		return 0, errors.New(fmt.Sprintf(`%s args.input must be slice`, `Min`))
	}
	for i := 0; i < inputRv.Len(); i++ {
		kind := inputRv.Index(i).Kind().String()
		var v float64
		if Includes(ReflectIntTypes, kind) {
			v = float64(inputRv.Index(i).Int())
		} else if Includes(ReflectFloatTypes, kind) {
			v = inputRv.Index(i).Float()
		} else {
			return 0, errors.New(fmt.Sprintf(`%s args.input element must be number`, `Min`))
		}
		if i == 0 {
			num = v
		}
		if num > v {
			num = v
		}
	}
	return num, nil
}

func Max(input interface{}) (num float64, err error) {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	if !Includes(ReflectArrayTypes, inputRv.Kind().String()) {
		return 0, errors.New(fmt.Sprintf(`%s args.input must be slice`, `Max`))
	}
	for i := 0; i < inputRv.Len(); i++ {
		kind := inputRv.Index(i).Kind().String()
		var v float64
		if Includes(ReflectIntTypes, kind) {
			v = float64(inputRv.Index(i).Int())
		} else if Includes(ReflectFloatTypes, kind) {
			v = inputRv.Index(i).Float()
		} else {
			return 0, errors.New(fmt.Sprintf(`%s args.input element must be number`, `Max`))
		}
		if i == 0 {
			num = v
		}
		if num < v {
			num = v
		}
	}
	return num, nil
}

func Sum(input interface{}) (num float64, err error) {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	if !Includes(ReflectArrayTypes, inputRv.Kind().String()) {
		return 0, errors.New(fmt.Sprintf(`%s args.input must be slice`, `Sum`))
	}
	decimalNum := decimal.NewFromFloat(0)
	for i := 0; i < inputRv.Len(); i++ {
		kind := inputRv.Index(i).Kind().String()
		var v decimal.Decimal
		if Includes(ReflectIntTypes, kind) {
			v = decimal.NewFromInt(inputRv.Index(i).Int())
		} else if kind == `float64` {
			v = decimal.NewFromFloat(inputRv.Index(i).Float())
		} else if kind == `float32` {
			return 0, errors.New(fmt.Sprintf(`%s not support float32, will loss precision...`, `Sum`))
		} else {
			return 0, errors.New(fmt.Sprintf(`%s args.input element must be number`, `Sum`))
		}
		decimalNum = decimalNum.Add(v)
	}
	num, _ = decimalNum.Float64()
	return num, nil
}

func Avg(input interface{}) (num float64, err error) {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	if !Includes(ReflectArrayTypes, inputRv.Kind().String()) {
		return 0, errors.New(fmt.Sprintf(`%s args.input must be slice`, `Avg`))
	}
	decimalCount := decimal.NewFromFloat(0)
	for i := 0; i < inputRv.Len(); i++ {
		kind := inputRv.Index(i).Kind().String()
		var v decimal.Decimal
		if Includes(ReflectIntTypes, kind) {
			v = decimal.NewFromInt(inputRv.Index(i).Int())
		} else if kind == `float64` {
			v = decimal.NewFromFloat(inputRv.Index(i).Float())
		} else if kind == `float32` {
			return 0, errors.New(fmt.Sprintf(`%s not support float32, will loss precision...`, `Avg`))
		} else {
			return 0, errors.New(fmt.Sprintf(`%s args.input element must be number`, `Avg`))
		}
		decimalCount = decimalCount.Add(v)
	}
	num, _ = decimalCount.Div(decimal.NewFromInt(int64(inputRv.Len()))).Float64()
	return num, nil
}

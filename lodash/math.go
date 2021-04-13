package lodash

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
)

// TODO
//  max  sum avg random [minBy maxBy sumBy]

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
		return 0, errors.New(fmt.Sprintf(`%s args.input must be slice`, `Max`))
	}
	decimalNum := decimal.NewFromFloat(0)
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
		decimalV := decimal.NewFromFloat(v)
		decimalNum = decimalNum.Add(decimalV)
	}
	num, _ = decimalNum.Float64()
	return num, nil
}

func Avg(input interface{}) (num float64, err error) {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	if !Includes(ReflectArrayTypes, inputRv.Kind().String()) {
		return 0, errors.New(fmt.Sprintf(`%s args.input must be slice`, `Max`))
	}
	decimalCount := decimal.NewFromFloat(0)
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
		decimalV := decimal.NewFromFloat(v)
		decimalCount = decimalCount.Add(decimalV)
	}
	count, _ := decimalCount.Float64()
	num, _ = decimal.NewFromFloat(count).Div(decimal.NewFromInt(int64(inputRv.Len()))).Float64()
	return num, nil
}

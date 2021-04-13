package lodash

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

func Chunk(output interface{}, input interface{}, sliceNum int) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	if sliceNum <= 0 {
		return errors.New(`Chunk args.sliceNum must > 0`)
	}
	err = CheckKindErr(`Chunk`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	groupNum := int(math.Ceil(float64(inputRv.Len()) / float64(sliceNum)))
	if isChain {
		result := []interface{}{}
		for i := 0; i < groupNum; i++ {
			endIndex := (i + 1) * sliceNum
			if endIndex > inputRv.Len() {
				endIndex = inputRv.Len()
			}
			result = append(result, inputRv.Slice((i*sliceNum), endIndex).Interface())
		}
		input.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for i := 0; i < groupNum; i++ {
			endIndex := (i + 1) * sliceNum
			if endIndex > inputRv.Len() {
				endIndex = inputRv.Len()
			}
			result = reflect.Append(result, inputRv.Slice((i*sliceNum), endIndex))
		}
		outputSet.Set(result)
	}
	return nil
}

func Concat(output interface{}, inputs ...interface{}) (err error) {
	_, isChain := chainArgConvert(output)
	if !isChain {
		if kind := reflect.ValueOf(output).Kind(); kind != reflect.Ptr {
			return errors.New(`Concat args.output must be pointer`)
		}
	}
	for _, input := range inputs {
		if kind := reflect.ValueOf(input).Kind(); kind != reflect.Slice {
			return errors.New(`Concat args.input must be slice`)
		}
	}

	if isChain {
		result := []interface{}{}
		for _, input := range inputs {
			inputReflect := reflect.ValueOf(input)
			for i := 0; i < inputReflect.Len(); i++ {
				result = append(result, inputReflect.Index(i).Interface())
			}
		}
		output.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for _, input := range inputs {
			inputReflect := reflect.ValueOf(input)
			for i := 0; i < inputReflect.Len(); i++ {
				result = reflect.Append(result, inputReflect.Index(i))
			}
		}
		outputSet.Set(result)
	}
	return nil
}

func Difference(output interface{}, input interface{}, accessory interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	if kind := reflect.ValueOf(accessory).Kind(); kind != reflect.Slice {
		return errors.New(`Difference args.accessory must be slice`)
	}
	err = CheckKindErr(`Difference`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	if isChain {
		result := []interface{}{}
		for i := 0; i < inputRv.Len(); i++ {
			if !Includes(accessory, inputRv.Index(i).Interface()) {
				result = append(result, inputRv.Index(i).Interface())
			}
		}
		input.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for i := 0; i < inputRv.Len(); i++ {
			if !Includes(accessory, inputRv.Index(i).Interface()) {
				result = reflect.Append(result, inputRv.Index(i))
			}
		}
		outputSet.Set(result)
	}
	return nil
}

func First(output interface{}, input interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	if length := inputRv.Len(); length == 0 {
		return errors.New(`First args.input is empty`)
	}
	err = CheckKindErr(`First`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	chainOutputNoSlice(output, input, isChain, inputRv.Index(0).Interface())
	return nil
}

func Last(output interface{}, input interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	if length := inputRv.Len(); length == 0 {
		return errors.New(`First args.input is empty`)
	}
	err = CheckKindErr(`Last`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	chainOutputNoSlice(output, input, isChain, inputRv.Index(inputRv.Len()-1).Interface())
	return nil
}

func flattenDepth(array interface{}, result *[]interface{}, level int) {
	level -= 1
	arrayRv := reflect.ValueOf(array)
	if !Includes(ReflectArrayTypes, arrayRv.Kind().String()) {
		arrayRv = arrayRv.Elem()
	}
	for i := 0; i < arrayRv.Len(); i++ {
		kind := arrayRv.Index(i).Kind().String()
		if kind == `interface` {
			kind = arrayRv.Index(i).Elem().Kind().String()
		}
		if level > 0 && Includes(ReflectArrayTypes, kind) {
			flattenDepth(arrayRv.Index(i).Interface(), result, level)
		} else {
			*result = append(*result, arrayRv.Index(i).Interface())
		}
	}
	return
}

func flattenForever(array interface{}, result *[]interface{}) {
	arrayRv := reflect.ValueOf(array)
	if !Includes(ReflectArrayTypes, arrayRv.Kind().String()) {
		arrayRv = arrayRv.Elem()
	}
	for i := 0; i < arrayRv.Len(); i++ {
		kind := arrayRv.Index(i).Kind().String()
		if kind == `interface` {
			kind = arrayRv.Index(i).Elem().Kind().String()
		}
		if Includes(ReflectArrayTypes, kind) {
			flattenForever(arrayRv.Index(i).Interface(), result)
		} else {
			*result = append(*result, arrayRv.Index(i).Interface())
		}
	}
	return
}

// if level <= 0, flatten forever.
func Flatten(output interface{}, input interface{}, level int) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`Flatten`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	result := []interface{}{}
	if level <= 0 {
		flattenForever(useInput, &result)
	} else {
		flattenDepth(useInput, &result, level+1)
	}
	err = chainOutputConvert(output, input, isChain, result)
	return err
}

func Uniq(output interface{}, input interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	err = CheckKindErr(`Uniq`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	if isChain {
		result := []interface{}{}
		for i := 0; i < inputRv.Len(); i++ {
			isExist := false
			if Includes(result, inputRv.Index(i).Interface()) {
				isExist = true
			}
			if !isExist {
				result = append(result, inputRv.Index(i).Interface())
			}
		}
		input.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for i := 0; i < inputRv.Len(); i++ {
			isExist := false
			if Includes(result.Interface(), inputRv.Index(i).Interface()) {
				isExist = true
			}
			if !isExist {
				result = reflect.Append(result, inputRv.Index(i))
			}
		}
		outputSet.Set(result)
	}
	return nil
}

func UniqBy(output interface{}, input interface{}, iteratee func(interface{}) interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	err = CheckKindErr(`UniqBy`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	if isChain {
		result := []interface{}{}
		for i := 0; i < inputRv.Len(); i++ {
			isExist := false
			for j := 0; j < len(result); j++ {
				if reflect.DeepEqual(iteratee(result[j]), iteratee(inputRv.Index(i).Interface())) {
					isExist = true
				}
			}
			if !isExist {
				result = append(result, inputRv.Index(i).Interface())
			}
		}
		input.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for i := 0; i < inputRv.Len(); i++ {
			isExist := false
			for j := 0; j < result.Len(); j++ {
				if reflect.DeepEqual(iteratee(result.Index(j).Interface()), iteratee(inputRv.Index(i).Interface())) {
					isExist = true
				}
			}
			if !isExist {
				result = reflect.Append(result, inputRv.Index(i))
			}
		}
		outputSet.Set(result)
	}
	return nil
}

func Union(output interface{}, inputs ...interface{}) {
	Chain(inputs[0]).Concat(inputs[1:]...).Uniq().Value(output)
}

func UnionBy(output interface{}, iteratee func(interface{}) interface{}, inputs ...interface{}) {
	Chain(inputs[0]).Concat(inputs[1:]...).UniqBy(iteratee).Value(output)
}

func IndexOf(input interface{}, iteratee func(interface{}) bool) int {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	index := -1
	for i := 0; i < inputRv.Len(); i++ {
		if iteratee(inputRv.Index(i).Interface()) {
			return i
		}
	}
	return index
}

func LastIndexOf(input interface{}, iteratee func(interface{}) bool) int {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	index := -1
	for i := 0; i < inputRv.Len(); i++ {
		if iteratee(inputRv.Index(i).Interface()) {
			index = i
		}
	}
	return index
}

func Join(output interface{}, input interface{}, joinStr string) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`Join`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	str := ``
	for i := 0; i < inputRv.Len(); i++ {
		kind := inputRv.Index(i).Kind().String()
		if Includes(ReflectIntTypes, kind) {
			str = str + strconv.Itoa(int(inputRv.Index(i).Int()))
		} else {
			str = str + inputRv.Index(i).String()
		}
		if i != inputRv.Len()-1 {
			str += joinStr
		}
	}
	chainOutputNoSlice(output, input, isChain, str)
	return nil
}

func Reverse(output interface{}, input interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`Reverse`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	if isChain {
		result := []interface{}{}
		for i := 1; i <= inputRv.Len(); i++ {
			result = append(result, inputRv.Index(inputRv.Len()-i).Interface())
		}
		input.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for i := 1; i <= inputRv.Len(); i++ {
			result = reflect.Append(result, inputRv.Index(inputRv.Len()-i))
		}
		outputSet.Set(result)
	}
	return nil
}

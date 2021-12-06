package lodash

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

//.Filter(func(user interface{}) bool {
//	return user.(User).Id == 3
//})
func Filter(output interface{}, input interface{}, iteratee func(interface{}) bool) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`Filter`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	if isChain {
		result := []interface{}{}
		for i := 0; i < inputRv.Len(); i++ {
			if iteratee(inputRv.Index(i).Interface()) {
				result = append(result, inputRv.Index(i).Interface())
			}
		}
		input.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for i := 0; i < inputRv.Len(); i++ {
			if iteratee(inputRv.Index(i).Interface()) {
				result = reflect.Append(result, inputRv.Index(i))
			}
		}
		outputSet.Set(result)
	}
	return nil
}

func Includes(input interface{}, checkValue interface{}) bool {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	switch inputRv.Kind().String() {
	case `array`, `slice`, `map`:
	default:
		panic(`Includes args.input must be slice!`)
	}

	isExist := false
	for i := 0; i < inputRv.Len(); i++ {
		if reflect.DeepEqual(inputRv.Index(i).Interface(), checkValue) {
			isExist = true
		}
	}
	return isExist
}

func IncludesBy(input interface{}, iteratee func(interface{}) bool) bool {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	switch inputRv.Kind().String() {
	case `array`, `slice`, `map`:
	default:
		panic(`IncludesBy args.input must be slice!`)
	}

	isExist := false
	for i := 0; i < inputRv.Len(); i++ {
		if iteratee(inputRv.Index(i).Interface()) {
			isExist = true
		}
	}
	return isExist
}

func Every(input interface{}, iteratee func(interface{}) bool) bool {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	switch inputRv.Kind().String() {
	case `array`, `slice`, `map`:
	default:
		panic(`Every args.input must be slice!`)
	}

	isAll := true
	for i := 0; i < inputRv.Len(); i++ {
		if !iteratee(inputRv.Index(i).Interface()) {
			isAll = false
		}
	}
	return isAll
}

func ForEach(input interface{}, iteratee func(interface{})) (err error) {
	useInput, _ := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	if kind := inputRv.Kind(); kind != reflect.Slice {
		return errors.New(`ForEach args.input must be slice!`)
	}

	for i := 0; i < inputRv.Len(); i++ {
		iteratee(inputRv.Index(i).Interface())
	}
	return nil
}

//.Map(func(user interface{}) interface{} {
//	newUser := User{}
//	reflect.ValueOf(&newUser).Elem().Set(reflect.ValueOf(user))
//	newUser.Name = `hahah`
//	return newUser.Id
//})

func Map(output interface{}, input interface{}, iteratee func(interface{}) interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`Map`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	if isChain {
		result := []interface{}{}
		for i := 0; i < inputRv.Len(); i++ {
			result = append(result, iteratee(inputRv.Index(i).Interface()))
		}
		input.(*lodash).input = result
	} else {
		outputSet := reflect.ValueOf(output).Elem()
		result := reflect.ValueOf(output).Elem()
		for i := 0; i < inputRv.Len(); i++ {
			result = reflect.Append(result, reflect.ValueOf(iteratee(inputRv.Index(i).Interface())))
		}
		outputSet.Set(result)
	}

	return nil
}

type GroupByItem struct {
	Key    interface{}
	Values []interface{}
}

type GroupByRes []GroupByItem

// output type must like []groupByObj.
func GroupBy(output interface{}, input interface{}, iteratee func(interface{}) (key interface{})) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`GroupBy`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	result := GroupByRes{}
	for i := 0; i < inputRv.Len(); i++ {
		groupKey := iteratee(inputRv.Index(i).Interface())
		groupIndex := IndexOf(result, func(v interface{}) bool {
			return v.(GroupByItem).Key == groupKey
		})
		if groupIndex == -1 {
			result = append(result, GroupByItem{
				Key:    groupKey,
				Values: []interface{}{inputRv.Index(i).Interface()},
			})
		} else {
			result[groupIndex].Values = append(result[groupIndex].Values, inputRv.Index(i).Interface())
		}
	}
	err = chainOutputConvert(output, input, isChain, result)
	return err
}

func Find(output interface{}, input interface{}, iteratee func(interface{}) bool) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	err = CheckKindErr(`Find`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	index := -1
	for i := 0; i < inputRv.Len(); i++ {
		if iteratee(inputRv.Index(i).Interface()) {
			index = i
		}
	}
	if index == -1 {
		return ErrNotFound
	}
	chainOutputNoSlice(output, input, isChain, inputRv.Index(index).Interface())
	return nil
}

type sorter struct {
	items []interface{}
	less  func(i, j interface{}) bool
}

func (s sorter) Len() int {
	return len(s.items)
}

func (s sorter) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s sorter) Less(i, j int) bool {
	return s.less(s.items[i], s.items[j])
}

func SortBy(output interface{}, input interface{}, iteratee func(interface{}) interface{}, order string) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`SortBy`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	items := []interface{}{}
	for i := 0; i < inputRv.Len(); i++ {
		items = append(items, inputRv.Index(i).Interface())
	}
	s := sorter{
		items: items,
		less: func(i, j interface{}) bool {
			i = iteratee(i)
			j = iteratee(j)
			valueKind := reflect.ValueOf(i).Kind().String()
			if Includes([]string{`array`, `slice`, `map`, `struct`, `ptr`, `chan`, `interface`, `func`}, valueKind) {
				panic(`SortBy compare value is not supported type!`)
			}
			if order == `asc` {
				if Includes(ReflectIntTypes, valueKind) {
					return reflect.ValueOf(i).Int() < reflect.ValueOf(j).Int()
				}
				if Includes(ReflectFloatTypes, valueKind) {
					return reflect.ValueOf(i).Float() < reflect.ValueOf(j).Float()
				}
				if valueKind == `string` {
					return reflect.ValueOf(i).String() < reflect.ValueOf(j).String()
				}
				if valueKind == `bool` {
					if reflect.ValueOf(i).Bool() == true {
						return false
					} else {
						return true
					}
				}
			}
			if order == `desc` {
				if Includes(ReflectIntTypes, valueKind) {
					return reflect.ValueOf(i).Int() > reflect.ValueOf(j).Int()
				}
				if Includes(ReflectFloatTypes, valueKind) {
					return reflect.ValueOf(i).Float() > reflect.ValueOf(j).Float()
				}
				if valueKind == `string` {
					return reflect.ValueOf(i).String() > reflect.ValueOf(j).String()
				}
				if valueKind == `bool` {
					if reflect.ValueOf(i).Bool() == true {
						return true
					} else {
						return false
					}
				}
			}
			return false
		},
	}
	sort.Sort(s)
	err = chainOutputConvert(output, input, isChain, s.items)
	return err
}

//iterateers := []func(interface{}) interface{}{
//	func(i interface{}) interface{} {
//		return i.(User).Id
//	},
//	func(i interface{}) interface{} {
//		return i.(User).Value
//	},
//	func(i interface{}) interface{} {
//		return i.(User).Name
//	},
//}
func OrderBy(output interface{}, input interface{}, iterateers []func(interface{}) interface{}, orders []string) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`SortBy`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	items := []interface{}{}
	for i := 0; i < inputRv.Len(); i++ {
		items = append(items, inputRv.Index(i).Interface())
	}
	s := sorter{
		items: items,
		less: func(i, j interface{}) bool {
			iResults := []interface{}{}
			jResults := []interface{}{}
			for _, iteratee := range iterateers {
				iResults = append(iResults, iteratee(i))
				jResults = append(jResults, iteratee(j))
			}
			for index := 0; index < len(iterateers); index++ {
				valueKind := reflect.ValueOf(iResults[index]).Kind().String()
				if Includes([]string{`array`, `slice`, `struct`, `chan`, `interface`, `func`}, valueKind) {
					panic(`SortBy compare value is not supported type!`)
				}
				// order default `asc`.
				order := `asc`
				if len(orders) >= index+1 {
					order = orders[index]
				}
				if order == `asc` {
					if Includes(ReflectIntTypes, valueKind) {
						left := fmt.Sprintf(`%d`, reflect.ValueOf(iResults[index]).Interface())
						right := fmt.Sprintf(`%d`, reflect.ValueOf(jResults[index]).Interface())
						if left < right {
							return true
						} else if left == right {
							continue
						} else {
							return false
						}
					}
					if Includes(ReflectFloatTypes, valueKind) {
						left := reflect.ValueOf(iResults[index]).Float()
						right := reflect.ValueOf(jResults[index]).Float()
						if left < right {
							return true
						} else if left == right {
							continue
						} else {
							return false
						}
					}
					if valueKind == `string` {
						left := reflect.ValueOf(iResults[index]).String()
						right := reflect.ValueOf(jResults[index]).String()
						if left < right {
							return true
						} else if left == right {
							continue
						} else {
							return false
						}
					}
					if valueKind == `bool` {
						left := reflect.ValueOf(iResults[index]).Bool()
						right := reflect.ValueOf(jResults[index]).Bool()
						if left == right {
							continue
						} else if left == false {
							return true
						} else {
							return false
						}
					}
				}
				if order == `desc` {
					if Includes(ReflectIntTypes, valueKind) {
						left := fmt.Sprintf(`%d`, reflect.ValueOf(iResults[index]).Interface())
						right := fmt.Sprintf(`%d`, reflect.ValueOf(jResults[index]).Interface())
						if left > right {
							return true
						} else if left == right {
							continue
						} else {
							return false
						}
					}
					if Includes(ReflectFloatTypes, valueKind) {
						left := reflect.ValueOf(iResults[index]).Float()
						right := reflect.ValueOf(jResults[index]).Float()
						if left > right {
							return true
						} else if left == right {
							continue
						} else {
							return false
						}
					}
					if valueKind == `string` {
						left := reflect.ValueOf(iResults[index]).String()
						right := reflect.ValueOf(jResults[index]).String()
						if left > right {
							return true
						} else if left == right {
							continue
						} else {
							return false
						}
					}
					if valueKind == `bool` {
						left := reflect.ValueOf(iResults[index]).Bool()
						right := reflect.ValueOf(jResults[index]).Bool()
						if left == right {
							continue
						} else if left == true {
							return true
						} else {
							return false
						}
					}
				}
			}
			return false
		},
	}
	sort.Sort(s)
	err = chainOutputConvert(output, input, isChain, s.items)
	return err
}

func Sort(output interface{}, input interface{}, key string, order string) (err error) {
	err = SortBy(output, input, func(i interface{}) interface{} {
		rv := reflect.ValueOf(i)
		if rv.Kind().String() == `struct` {
			return reflect.ValueOf(i).FieldByName(key).Interface()
		} else {
			return i
		}
	}, order)
	return
}

func wrapOrder(iterateers *[]func(interface{}) interface{}, key string) {
	*iterateers = append(*iterateers, func(i interface{}) interface{} {
		rv := reflect.ValueOf(i)
		switch rv.Kind().String() {
		case `struct`:
			return rv.FieldByName(key).Interface()
		case `map`:
			newV := reflect.ValueOf(i).Interface().(map[string]interface{})
			return newV[key]
		case `ptr`:
			newV := rv.Elem().Interface()
			return reflect.ValueOf(newV).FieldByName(key).Interface()
		default:
			return i
		}
	})
}

func Order(output interface{}, input interface{}, keys []string, orders []string) (err error) {
	iterateers := []func(interface{}) interface{}{}
	for _, key := range keys {
		wrapOrder(&iterateers, key)
	}
	err = OrderBy(output, input, iterateers, orders)
	return
}

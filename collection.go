package lodash

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

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
//	newUser.Name = `tom`
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

func SortBy(output interface{}, input interface{}, valueFunction func(interface{}) interface{}, order string) (err error) {
	return OrderBy(output, input, []func(interface{}) interface{}{valueFunction}, []string{order})
}

type LessFunc func(p, q interface{}) bool
type multiSorter struct {
	changes []interface{}
	less    []LessFunc
}

func OrderByLess(less ...LessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}
func (ms *multiSorter) Sort(changes []interface{}) {
	ms.changes = changes
	sort.Sort(ms)
}

func (ms *multiSorter) Len() int {
	return len(ms.changes)
}
func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}
func (ms *multiSorter) Less(i, j int) bool {
	p, q := ms.changes[i], ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func OrderBy(output interface{}, input interface{}, valueFunctions []func(interface{}) interface{}, orders []string) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)
	err = CheckKindErr(`OrderBy`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}

	items := []interface{}{}
	for i := 0; i < inputRv.Len(); i++ {
		items = append(items, inputRv.Index(i).Interface())
	}

	var lessFunctions []LessFunc
	for index := range valueFunctions {
		func(i int) {
			lessFunctions = append(lessFunctions, func(p, q interface{}) bool {
				// order default `asc`.
				order := `asc`
				if len(orders) >= i+1 && orders[i] == `desc` {
					order = `desc`
				}
				pFuncValue := valueFunctions[i](p)
				qFuncValue := valueFunctions[i](q)
				pRv := reflect.ValueOf(pFuncValue)
				qRv := reflect.ValueOf(qFuncValue)

				valueKind := pRv.Kind().String()
				if Includes([]string{`array`, `slice`, `struct`, `chan`, `interface`, `func`}, valueKind) {
					panic(`OrderBy compare value is not supported type!`)
				}
				if Includes(ReflectFloatTypes, valueKind) {
					if pRv.Float() == qRv.Float() {
						return false
					}
					if order == `asc` {
						return pRv.Float() < qRv.Float()
					} else {
						return pRv.Float() > qRv.Float()
					}
				}

				pValue := fmt.Sprintf(`%v`, pFuncValue)
				qValue := fmt.Sprintf(`%v`, qFuncValue)
				// bool to string: "0" / "1"
				if valueKind == `bool` {
					pValue, qValue = `0`, `0`
					if pRv.Bool() {
						pValue = `1`
					}
					if qRv.Bool() {
						qValue = `1`
					}
				}

				// "equal" == "not less", next less().
				if pValue == qValue {
					return false
				}
				if order == `asc` {
					return pValue < qValue
				} else {
					return pValue > qValue
				}
			})
		}(index)
	}
	OrderByLess(lessFunctions...).Sort(items)
	err = chainOutputConvert(output, input, isChain, items)
	return err
}

func Sort(output interface{}, input interface{}, key string, order string) (err error) {
	var valueFunctions []func(interface{}) interface{}
	wrapOrder(&valueFunctions, key)
	err = OrderBy(output, input, valueFunctions, []string{order})
	return
}

func wrapOrder(valueFunctions *[]func(interface{}) interface{}, key string) {
	*valueFunctions = append(*valueFunctions, func(i interface{}) interface{} {
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
	var valueFunctions []func(interface{}) interface{}
	for _, key := range keys {
		wrapOrder(&valueFunctions, key)
	}
	err = OrderBy(output, input, valueFunctions, orders)
	return
}

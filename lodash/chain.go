package lodash

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type lodash struct {
	input interface{}
	err   error
}

func Chain(input interface{}) *lodash {
	l := lodash{}
	l.input = input
	return &l
}

func (l *lodash) Value(output interface{}) error {
	if l.err != nil {
		return l.err
	}
	// 简单类型和struct可以直接反射set
	inputKind := reflect.ValueOf(l.input).Kind().String()
	switch inputKind {
	case `array`, `map`, `slice`:
		outputJson, err := json.Marshal(l.input)
		if err != nil {
			return errors.New(fmt.Sprintf(`output format error: %s`, err.Error()))
		}
		err = json.Unmarshal(outputJson, output)
		if err != nil {
			return errors.New(fmt.Sprintf(`output format error: %s`, err.Error()))
		}
	default:
		reflect.ValueOf(output).Elem().Set(reflect.ValueOf(l.input))
	}
	return nil
}

func (l *lodash) Chunk(sliceNum int) *lodash {
	if l.err != nil {
		return l
	}
	err := Chunk(nil, l, sliceNum)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Concat(inputs ...interface{}) *lodash {
	if l.err != nil {
		return l
	}
	newInputs := []interface{}{l.input}
	for _, input := range inputs {
		newInputs = append(newInputs, input)
	}
	err := Concat(l, newInputs...)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Difference(accessory interface{}) *lodash {
	if l.err != nil {
		return l
	}
	err := Difference(nil, l, accessory)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Uniq() *lodash {
	if l.err != nil {
		return l
	}
	err := Uniq(nil, l)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) UniqBy(iteratee func(interface{}) interface{}) *lodash {
	if l.err != nil {
		return l
	}
	err := UniqBy(nil, l, iteratee)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Union(inputs ...interface{}) *lodash {
	if l.err != nil {
		return l
	}
	l.Concat(inputs...).Uniq()
	return l
}

func (l *lodash) UnionBy(iteratee func(interface{}) interface{}, inputs ...interface{}) *lodash {
	if l.err != nil {
		return l
	}
	l.Concat(inputs...).UniqBy(iteratee)
	return l
}

func (l *lodash) Filter(iteratee func(interface{}) bool) *lodash {
	if l.err != nil {
		return l
	}
	err := Filter(nil, l, iteratee)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Includes(checkValue interface{}) *lodash {
	if l.err != nil {
		return l
	}
	l.input = Includes(l, checkValue)
	return l
}

func (l *lodash) IncludesBy(iteratee func(interface{}) bool) *lodash {
	if l.err != nil {
		return l
	}
	l.input = IncludesBy(l, iteratee)
	return l
}

func (l *lodash) Every(iteratee func(interface{}) bool) *lodash {
	if l.err != nil {
		return l
	}
	l.input = Every(l, iteratee)
	return l
}

func (l *lodash) ForEach(iteratee func(interface{})) *lodash {
	if l.err != nil {
		return l
	}
	err := ForEach(l, iteratee)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Map(iteratee func(interface{}) interface{}) *lodash {
	if l.err != nil {
		return l
	}
	err := Map(nil, l, iteratee)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) GroupBy(iteratee func(interface{}) (key interface{})) *lodash {
	if l.err != nil {
		return l
	}
	err := GroupBy(nil, l, iteratee)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Find(iteratee func(interface{}) bool) *lodash {
	if l.err != nil {
		return l
	}
	err := Find(nil, l, iteratee)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Reverse() *lodash {
	if l.err != nil {
		return l
	}
	err := Reverse(nil, l)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Join(joinStr string) *lodash {
	if l.err != nil {
		return l
	}
	err := Join(nil, l, joinStr)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) IndexOf(iteratee func(interface{}) bool) *lodash {
	if l.err != nil {
		return l
	}
	l.input = IndexOf(l, iteratee)
	return l
}

func (l *lodash) LastIndexOf(iteratee func(interface{}) bool) *lodash {
	if l.err != nil {
		return l
	}
	l.input = LastIndexOf(l, iteratee)
	return l
}

func (l *lodash) First() *lodash {
	if l.err != nil {
		return l
	}
	err := First(nil, l)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Last() *lodash {
	if l.err != nil {
		return l
	}
	err := Last(nil, l)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Flatten(level int) *lodash {
	if l.err != nil {
		return l
	}
	err := Flatten(nil, l, level)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) SortBy(iteratee func(interface{}) interface{}, order string) *lodash {
	if l.err != nil {
		return l
	}
	err := SortBy(nil, l, iteratee, order)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) OrderBy(iterateers []func(interface{}) interface{}, orders []string) *lodash {
	if l.err != nil {
		return l
	}
	err := OrderBy(nil, l, iterateers, orders)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Sort(key string, order string) *lodash {
	if l.err != nil {
		return l
	}
	err := Sort(nil, l, key, order)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) Order(keys []string, orders []string) *lodash {
	if l.err != nil {
		return l
	}
	err := Order(nil, l, keys, orders)
	if err != nil {
		l.err = err
		return l
	}
	return l
}

func (l *lodash) ConcatStr(inputs ...string) *lodash {
	if l.err != nil {
		return l
	}
	newStr, err := ConcatStr(l.input, inputs...)
	if err != nil {
		l.err = err
		return l
	}
	l.input = newStr
	return l
}

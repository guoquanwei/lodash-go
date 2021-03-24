## Introduction
lodash-go is goLang methods package, it like javascript's lodash.

### go doc reference
[![GoDoc](https://godoc.org/github.com/ITcathyh/alloter?status.svg)](https://godoc.org/github.com/guoquanwei/lodash-go)

### Variables

```go
func Chain(input interface{}) *lodash
// init a chain
// example: Chain([]some_struct)

func Value(output interface{}) error
// output and bind result
// example: Chain([]some_struct{...}).some_methods().Value(&result)

func Chunk(output interface{}, input interface{}, sliceNum int) (err error)

func Concat(output interface{}, inputs ...interface{}) (err error)
// support concat multi arrays
// example: Concat(&result, []some_struct{...}, []some_struct{...})

func Difference(output interface{}, input interface{}, accessory interface{}) (err error)
// if `input` not includes `accessory`'s element， will append to `input`.
// example: Difference(&result, []some_struct{...}, []some_struct{...})

func Includes(input interface{}, checkValue interface{}) bool
// check `input` contains `checkValue`.

func IncludesBy(input interface{}, iteratee func(interface{}) bool) bool
// check `input` contains match condition value.

func Filter(output interface{}, input interface{}, iteratee func(interface{}) bool) (err error)
// all match condition elements append to result.
// example: Filter(&result,[]some_struct{...}, func(v interface{}) bool { return v.(some_struct).Id > 1 })

func Every(input interface{}, iteratee func(interface{}) bool) bool
// if all `input`'s element can be true(iteratee(element)), will return true.
// example: Every([]some_struct{...}, func(v interface{}) bool { return v.(some_struct).Id > 1 })

func Find(output interface{}, input interface{}, iteratee func(interface{}) bool) (err error)
// return first match condition element.

func IndexOf(input interface{}, iteratee func(interface{}) bool) int
// return first match condition element's index.

func LastIndexOf(input interface{}, iteratee func(interface{}) bool) int
// return last match condition element's index.

func First(output interface{}, input interface{}) (err error)
// return first element. like array[0]

func Last(output interface{}, input interface{}) (err error)
// return last element

func Flatten(output interface{}, input interface{}, level int) (err error)
// level is flatten total, if level <= 0, will flatten forever.

func ForEach(input interface{}, iteratee func(interface{})) (err error)
// just a iterateer.

func Map(output interface{}, input interface{}, iteratee func(interface{}) interface{}) (err error)
// iteratee `input`'s elements and return new value.
// example: Map(&result, []some_struct{...}, func(v interface{}) interface{} {
//   newStruct := some_Struct{}
//   newStruct.Id = v.(some_Struct).Id + 1
//   return newStruct
// })

func GroupBy(output interface{}, input interface{}, iteratee func(interface{}) (key interface{})) (err error)
// group by iteratee function's return value.
// example:
// type someGroup struct {
//   Key string `json:"key"`
//   Values []some_struct `json:"values"`
// }
// groups := []someGroup{}
// lodash.GroupBy(&groups, []some_struct{...}, func(i interface{}) (key interface{}) {
//   return i.(some_struct).Id
// })

func Order(output interface{}, input interface{}, keys []string, orders []string) (err error)
// example: Order(&result, []some_struct{...}, []string{`k1`, `k2`}, []string{`asc`, `desc`})
// if len(orders) < len(keys), default `asc`.

func OrderBy(output interface{}, input interface{}, iterateers []func(interface{}) interface{}, orders []string) (err error)
// like Order, iterateers funtion's return values is keys.

func Sort(output interface{}, input interface{}, key string, order string) (err error)
// like Order, but just support one key.

func SortBy(output interface{}, input interface{}, iteratee func(interface{}) interface{}, order string) (err error)
// like Sort, iterate funtion's return value is key.

func Reverse(output interface{}, input interface{}) (err error)
// reverse `input` array.

func Uniq(output interface{}, input interface{}) (err error)
// return Unique elements.

func UniqBy(output interface{}, input interface{}, iteratee func(interface{}) interface{}) (err error)
// return Unique elements by function's return value.

func Union(output interface{}, inputs ...interface{})
// like Concat + Uniq.

func UnionBy(output interface{}, iteratee func(interface{}) interface{}, inputs... interface{})
// like Concat + UniqBy.

func Join(output interface{}, input interface{}, joinStr string) (err error)
// array to string, each element use `joinStr` link.

```

### Demo
```go
package main

import "github.com/guoquanwei/lodash-go"

type User struct {
	Id   int
	Name string
}

func main ()  {
	// simple method
	newUsers := []User{}
	err1 := lodash.Filter(&newUsers, []User{{Id: 1}, {Id: 2}, func(i interface{}) bool {
		return i.(User).Id > 1
	}})

	// multi methods use lodash.Chain()
	user := User{}
	err2 := lodash.Chain([]User{{Id: 1}, {Id: 2}).Filter(func(i interface{}) bool {
		return i.(User).Id > 1
	}}).First().Value(&user)
	
}

```

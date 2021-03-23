## Introduction
lodash-go is goLang methods package, it like javascript's lodash.

### API doc reference
[![GoDoc](https://godoc.org/github.com/ITcathyh/alloter?status.svg)](https://godoc.org/github.com/guoquanwei/lodash-go)

### Variables

```
func Chain(input interface{}) *lodash

func CheckKindErr(funcName string, isChain bool, outputKind string, inputKind string) error

func Chunk(output interface{}, input interface{}, sliceNum int) (err error)

func Concat(output interface{}, inputs ...interface{}) (err error)

func Difference(output interface{}, input interface{}, accessory interface{}) (err error)

func Every(input interface{}, iteratee func(interface{}) bool) bool

func Filter(output interface{}, input interface{}, iteratee func(interface{}) bool) (err error)

func Find(output interface{}, input interface{}, iteratee func(interface{}) bool) (err error)

func First(output interface{}, input interface{}) (err error)

func Flatten(output interface{}, input interface{}, level int) (err error)

func ForEach(input interface{}, iteratee func(interface{})) (err error)

func GroupBy(output interface{}, input interface{}, ...) (err error)

func Includes(input interface{}, checkValue interface{}) bool

func IncludesBy(input interface{}, iteratee func(interface{}) bool) bool

func IndexOf(input interface{}, iteratee func(interface{}) bool) int

func Join(output interface{}, input interface{}, joinStr string) (err error)

func Last(output interface{}, input interface{}) (err error)

func LastIndexOf(input interface{}, iteratee func(interface{}) bool) int

func Map(output interface{}, input interface{}, iteratee func(interface{}) interface{}) (err error)

func Order(output interface{}, input interface{}, keys []string, orders []string) (err error)

func OrderBy(output interface{}, input interface{}, ...) (err error)

func Reverse(output interface{}, input interface{}) (err error)

func Sort(output interface{}, input interface{}, key string, order string) (err error)

func SortBy(output interface{}, input interface{}, iteratee func(interface{}) interface{}, ...) (err error)

func Union(output interface{}, inputs ...interface{})

func UnionBy(output interface{}, iteratee func(interface{}) interface{}, ...)

func Uniq(output interface{}, input interface{}) (err error)

func UniqBy(output interface{}, input interface{}, iteratee func(interface{}) interface{}) (err error)

```


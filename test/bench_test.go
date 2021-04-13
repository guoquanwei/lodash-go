package test

import (
	"encoding/json"
	"reflect"
	"testing"
)

type User struct {
	Id int
}

func Benchmark_Tran_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newUsers := User{}
		rv, _ := json.Marshal(User{Id: 1})
		json.Unmarshal(rv, &newUsers)
	}

}

func Benchmark_Tran_Reflect (b *testing.B)  {
	for i := 0; i < b.N; i++ {
		newUsers := User{}
		reflect.ValueOf(&newUsers).Elem().Set(reflect.ValueOf(User{Id: 1}))
	}
}
package util

import (
	"fmt"
	"reflect"
)

func SliceContains(slice interface{}, value interface{}) bool {
	sType := reflect.TypeOf(slice)
	vType := reflect.TypeOf(value)
	if sType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("SliceContains passed non slice %T", slice))
	}

	if sType.Elem().Kind() != vType.Kind() {
		panic(fmt.Sprintf("SliceContains passed non-matching slice and value types %T, %T", slice, value))
	}

	rSlice := reflect.ValueOf(slice)
	for i := 0; i < rSlice.Len(); i++ {
		v := rSlice.Index(i)
		if reflect.DeepEqual(v.Interface(), value) {
			return true
		}
	}
	return false
}

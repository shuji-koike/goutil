package goutil

import "reflect"

// SliceMap returns a slice with the results of calling a provided function on every element in the given slice.
// SliceMap([]T, func(T) V) []V
func SliceMap(slice interface{}, fn interface{}) interface{} {
	sliceV := reflect.ValueOf(slice)
	fnV := reflect.ValueOf(fn)
	size := sliceV.Len()
	ret := reflect.MakeSlice(reflect.SliceOf(fnV.Type().Out(0)), size, size)
	for i := 0; i < size; i++ {
		ret.Index(i).Set(fnV.Call([]reflect.Value{sliceV.Index(i)})[0])
	}
	return ret.Interface()
}

// SliceReduce returns a accumulated value.
// SliceReduce([]T, V, func(V, T) V) V
func SliceReduce(slice interface{}, acc interface{}, fn interface{}) interface{} {
	sliceV := reflect.ValueOf(slice)
	fnV := reflect.ValueOf(fn)
	accV := reflect.ValueOf(acc)
	size := sliceV.Len()
	for i := 0; i < size; i++ {
		accV = (fnV.Call([]reflect.Value{accV, sliceV.Index(i)})[0])
	}
	return accV.Interface()
}

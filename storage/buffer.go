package storage

import "reflect"
import "unsafe"

type Buffer interface {
	Copy(i, j int)
	Swap(i, j int)
	Clear(i, n int)
}

var _BYTE = reflect.Typeof(byte(0))
var _BYTE_SLICE = reflect.Typeof([]byte{})
var _SLICE_HEADER = reflect.Typeof(reflect.SliceHeader{})

var _UINT = reflect.Typeof(uint(0))
var _UINT_SIZE = unsafe.Sizeof(uint(0))
var _UINT_BUFFER = reflect.Typeof([]uint{})

var _PTR = reflect.Typeof(float(0))
var _PTR_SIZE = unsafe.Sizeof(uintptr(0))
var _PTR_BUFFER = reflect.Typeof([]uintptr{})


func AsByteSlice(b interface{}) []byte {
	if v, ok := reflect.NewValue(b).(*reflect.SliceValue); ok {
		typesize := int(v.Type().(*reflect.SliceType).Elem().Size())
		h := unsafe.Unreflect(_SLICE_HEADER, unsafe.Pointer(v.Addr())).(reflect.SliceHeader)
		h.Len = typesize * v.Len()
		h.Cap = typesize * v.Cap()
	 	return unsafe.Unreflect(_BYTE_SLICE, unsafe.Pointer(&h)).([]byte)
	}
	panic(b)
}

func AsFloatBuffer(b interface{}) FloatBuffer {
	if v, ok := reflect.NewValue(b).(*reflect.SliceValue); ok {
		typesize := int(v.Type().(*reflect.SliceType).Elem().Size())
		h := unsafe.Unreflect(_SLICE_HEADER, unsafe.Pointer(v.Addr())).(reflect.SliceHeader)
		h.Len = typesize * v.Len() / _FLOAT_SIZE
		h.Cap = typesize * v.Cap() / _FLOAT_SIZE
 		return unsafe.Unreflect(_FLOAT_BUFFER, unsafe.Pointer(&h)).(FloatBuffer)
	}
	panic(b)
}

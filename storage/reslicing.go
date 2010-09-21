package storage

import "reflect"
import "unsafe"

type Resliceable interface {
	SliceHeader(element_size int) *reflect.SliceHeader
}

var _BYTE = reflect.Typeof(byte(0))
var _BYTE_SLICE = reflect.Typeof([]byte(nil))

func AsByteSlice(b Resliceable) []byte {
	return unsafe.Unreflect(_BYTE_SLICE, unsafe.Pointer(b.SliceHeader(1))).([]byte)
}

var _INT_SIZE = unsafe.Sizeof(int(0))
var _INT_BUFFER = reflect.Typeof(IntBuffer(nil))

func AsIntBuffer(b Resliceable) IntBuffer {
	return unsafe.Unreflect(_INT_BUFFER, unsafe.Pointer(b.SliceHeader(_INT_SIZE))).(IntBuffer)
}

var _UINT_SIZE = unsafe.Sizeof(uint(0))
var _UINT_BUFFER = reflect.Typeof([]uint{})

func AsUintBuffer(b Resliceable) []uint {
	return unsafe.Unreflect(_UINT_BUFFER, unsafe.Pointer(b.SliceHeader(_UINT_SIZE))).([]uint)
}

var _FLOAT_SIZE = unsafe.Sizeof(float(0))
var _FLOAT_BUFFER = reflect.Typeof(FloatBuffer{})

func AsFloatBuffer(b Resliceable) FloatBuffer {
	return unsafe.Unreflect(_FLOAT_BUFFER, unsafe.Pointer(b.SliceHeader(_FLOAT_SIZE))).(FloatBuffer)
}

var _PTR_SIZE = unsafe.Sizeof(uintptr(0))
var _PTR_BUFFER = reflect.Typeof([]uintptr{})

func AsPointerBuffer(b Resliceable) []uintptr {
	return unsafe.Unreflect(_PTR_BUFFER, unsafe.Pointer(b.SliceHeader(_PTR_SIZE))).([]uintptr)
}

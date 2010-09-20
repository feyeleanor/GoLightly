package storage

import "reflect"
import "unsafe"

type IntBuffer []int

func (b IntBuffer) Len() int							{ return len(b) }
func (b IntBuffer) Cap() int							{ return cap(b) }
func (b IntBuffer) Copy(i, j int)						{ b[i] = b[j] }
func (b IntBuffer) Swap(i, j int)						{ b[i], b[j] = b[j], b[i] }
func (b IntBuffer) Clear(i, j int) {
	if j > len(b) {
		j = len(b)
	}
	for ; i < j; i++ {
		b[i] = 0
	}
}

func (b IntBuffer) Clone() IntBuffer {
	s := make(IntBuffer, len(b))
	copy(s, b)
	return s
}

func (b IntBuffer) Replicate(count int) IntBuffer {
	l := len(b)
	s := make(IntBuffer, l * count)
	for i, offset := 0, 0; i < count; i++ {
		copy(s[offset:], b)
		offset += l
	}
	return s
}

func (b *IntBuffer) Resize(length int) {
	if length > cap(*b) {
		x := *b
		*b = make(IntBuffer, length)
		copy(*b, x)
	} else {
		*b = (*b)[:length]
	}
}

func (b *IntBuffer) Extend(count int) {
	b.Resize(len(*b) + count)
}

func (b *IntBuffer) Shrink(count int) {
	b.Resize(len(*b) - count)
}

func (b IntBuffer) Collect(f func(i int, x int) int) IntBuffer {
	n := make(IntBuffer, len(b))
	for i, x := range b {
		n[i] = f(i, x)
	}
	return n
}

func (b IntBuffer) Negate(i int)						{ b[i] = -b[i] }
func (b IntBuffer) Increment(i int)						{ b[i] += 1 }
func (b IntBuffer) Decrement(i int)						{ b[i] -= 1 }

func (b IntBuffer) Add(i, j int)						{ b[i] += b[j] }
func (b IntBuffer) Subtract(i, j int)					{ b[i] -= b[j] }
func (b IntBuffer) Multiply(i, j int)					{ b[i] *= b[j] }
func (b IntBuffer) Divide(i, j int)						{ b[i] /= b[j] }
func (b IntBuffer) Remainder(i, j int)					{ b[i] %= b[j] }

func (b IntBuffer) And(i, j int)						{ b[i] &= b[j] }
func (b IntBuffer) Or(i, j int)							{ b[i] |= b[j] }
func (b IntBuffer) Xor(i, j int)						{ b[i] ^= b[j] }
func (b IntBuffer) Invert(i int)						{ b[i] = ^b[i] }
func (b IntBuffer) ShiftLeft(i, j int)					{ b[i] <<= uint(b[j]) }
func (b IntBuffer) ShiftRight(i, j int)					{ b[i] >>= uint(b[j]) }

func (b IntBuffer) Less(i, j int) bool					{ return b[i] < b[j] }
func (b IntBuffer) Equal(i, j int) bool					{ return b[i] == b[j] }
func (b IntBuffer) Greater(i, j int) bool				{ return b[i] > b[j] }
func (b IntBuffer) ZeroLess(i int) bool					{ return b[i] < 0 }
func (b IntBuffer) ZeroEqual(i int) bool				{ return b[i] == 0 }
func (b IntBuffer) ZeroGreater(i int) bool				{ return b[i] > 0 }


var _INT = reflect.Typeof(int(0))
var _INT_SIZE = unsafe.Sizeof(int(0))
var _INT_BUFFER = reflect.Typeof(IntBuffer{})

func AsIntBuffer(b interface{}) IntBuffer {
	if v, ok := reflect.NewValue(b).(*reflect.SliceValue); ok {
		typesize := int(v.Type().(*reflect.SliceType).Elem().Size())
		h := unsafe.Unreflect(_SLICE_HEADER, unsafe.Pointer(v.Addr())).(reflect.SliceHeader)
		h.Len = typesize * v.Len() / _INT_SIZE
		h.Cap = typesize * v.Cap() / _INT_SIZE
 		return unsafe.Unreflect(_INT_BUFFER, unsafe.Pointer(&h)).(IntBuffer)
	}
	panic(b)
}

func (b IntBuffer) ByteSlice() []byte {
	h := unsafe.Unreflect(_SLICE_HEADER, unsafe.Pointer(&b)).(reflect.SliceHeader)
	h.Len = _INT_SIZE * len(b)
	h.Cap = _INT_SIZE * cap(b)
 	return unsafe.Unreflect(_BYTE_SLICE, unsafe.Pointer(&h)).([]byte)
}

func (b IntBuffer) FloatBuffer() FloatBuffer {
	h := unsafe.Unreflect(_SLICE_HEADER, unsafe.Pointer(&b)).(reflect.SliceHeader)
	h.Len = _INT_SIZE * len(b) / _FLOAT_SIZE
	h.Cap = _INT_SIZE * cap(b) / _FLOAT_SIZE
 	return unsafe.Unreflect(_FLOAT_BUFFER, unsafe.Pointer(&h)).(FloatBuffer)
}

func (b IntBuffer) Feed(c chan<- int, f func(i, x int) int) {
	d := b.Clone()
	go func() {
		for i, v := range d { c <- f(i, v) }
		close(c)
	}()
}

func (b IntBuffer) Pipe(f func(i, x int) int) <-chan int {
	c := make(chan int)
	b.Feed(c, f)
	return c
}

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

func (b IntBuffer) Collect(f func(x int) int) IntBuffer {
	n := make(IntBuffer, len(b))
	for i, x := range b {
		n[i] = f(x)
	}
	return n
}

func (b IntBuffer) Inject(seed int, f func(memo, x int) int) int {
	for _, x := range b {
		seed = f(seed, x)
	}
	return seed
}

func (b IntBuffer) Cycle(count int, f func(i, x int)) (j int) {
	switch count {
	case 0:		for {
					for _, x := range b {
						f(j, x)
					}
					j++
				}
	default:	for k := 0; j < count; j++ {
					for _, x := range b {
						f(k, x)
					}
					k++
				}
	}
	return
}

func (b IntBuffer) Combine(o IntBuffer, f func(x, y int) int) IntBuffer {
	if len(b) != len(o) {
		panic(b)
	}
	n := make(IntBuffer, len(b))
	for i, x := range b {
		n[i] = f(x, o[i])
	}
	return n
}

func (b IntBuffer) Count(f func(x int) bool) (c int) {
	for _, v := range b {
		if f(v) {
			c++
		}
	}
	return
}

func (b IntBuffer) Any(f func(x int) bool) bool {
	for _, v := range b {
		if f(v) {
			return true
		}
	}
	return false
}

func (b IntBuffer) All(f func(x int) bool) bool {
	for _, v := range b {
		if !f(v) {
			return false
		}
	}
	return true
}

func (b IntBuffer) None(f func(x int) bool) bool {
	for _, v := range b {
		if f(v) {
			return false
		}
	}
	return true
}

func (b IntBuffer) One(f func(x int) bool) bool {
	c := 0
	for _, v := range b {
		switch {
		case c > 1:		return false
		case f(v):		c++
		}
	}
	return c == 1
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


func (b IntBuffer) SliceHeader(element_size int) *reflect.SliceHeader {
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Len = int(float(len(b)) * float(_INT_SIZE / element_size))
	h.Cap = int(float(cap(b)) * float(_INT_SIZE / element_size))
	return &h
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

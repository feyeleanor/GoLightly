//	TODO:	Storing and retrieving of pointers
//	TODO:	Shouldn't support Extend or Expand
//	TODO:	Interface?

package vm

import "unsafe"

type Buffer []int

func (b *Buffer) iterate(c chan<- int) {
	for _, v := range *b { c <- v }
	close(c)
}

func (b *Buffer) Iter() <-chan int {
	c := make(chan int)
	go b.iterate(c)
	return c
}

func (b *Buffer) realloc(length int) (a []int) {
	a = make(Buffer, length)
	copy(a, *b)
	*b = a
	return
}

func (b *Buffer) Extend(count int) {
	a := *b
	desired_length := len(a) + count
	if desired_length > cap(a) {
		a = b.realloc(desired_length)
	} else {
		a = a[0:desired_length - 1]
	}
	*b = a
}

func (b *Buffer) Resize(length int) *Buffer {
	a := *b
	if length > cap(a) {
		a = b.realloc(length)
	} else if length < len(a) {
		for i := range a[length:] {
			var zero int
			a[length+i] = zero
		}
	}
	*b = a[0:length]
	return b
}

func (b *Buffer) Slice(i, j int) *Buffer {
	var s Buffer
	if (j - i) < 0 {
		s.realloc(0)
	} else {
		s.realloc(j - i)
		copy(s, (*b)[i:j])
	}
	return &s
}

func (b *Buffer) Init(length int)						{ b.Resize(length) }
func (b *Buffer) Len() int								{ return len(*b) }
func (b *Buffer) Cap() int								{ return cap(*b) }
func (b *Buffer) At(i int) int							{ return (*b)[i] }
func (b *Buffer) Set(i int, x int)						{ (*b)[i] = x }
func (b *Buffer) FAt(i int) float						{ return *(*float)(unsafe.Pointer(&(*b)[i])) }
func (b *Buffer) FSet(i int, f float)					{ (*b)[i] = *(*int)(unsafe.Pointer(&f)) }

func (b *Buffer) First() int							{ return (*b)[0] }
func (b *Buffer) Last() int								{ return (*b)[len(*b)-1] }
func (b *Buffer) Clone() *Buffer						{ return b.Slice(0, b.Len()) }
func (b *Buffer) Replace(o *Buffer)						{ copy(*b, *o) }
func (b *Buffer) Identical(o *Buffer) bool {
	match := b.Len() == o.Len()
	if match { for i, e := range *o { match = match && b.At(i) == e } }
	return match
}

func (b *Buffer) Do(f func(element int))				{ for _, e := range *b { f(e) } }

func (b *Buffer) Add(i, j int)							{ (*b)[i] += (*b)[j] }
func (b *Buffer) Subtract(i, j int)						{ (*b)[i] -= (*b)[j] }
func (b *Buffer) Multiply(i, j int)						{ (*b)[i] *= (*b)[j] }
func (b *Buffer) Divide(i, j int)						{ (*b)[i] /= (*b)[j] }
func (b *Buffer) FAdd(i, j int)							{ b.FSet(i, b.FAt(i) + b.FAt(j)) }
//func (b *Buffer) FSubtract(i, j int)					{ (*b)[i] -= (*b)[j] }
//func (b *Buffer) FMultiply(i, j int)					{ (*b)[i] *= (*b)[j] }
//func (b *Buffer) FDivide(i, j int)						{ (*b)[i] /= (*b)[j] }
func (b *Buffer) And(i, j int)							{ (*b)[i] &= (*b)[j] }
func (b *Buffer) Or(i, j int)							{ (*b)[i] |= (*b)[j] }
func (b *Buffer) Xor(i, j int)							{ (*b)[i] ^= (*b)[j] }
func (b *Buffer) Clear(i int)							{ (*b)[i] = 0 }
func (b *Buffer) Increment(i int)						{ (*b)[i] += 1 }
func (b *Buffer) Decrement(i int)						{ (*b)[i] -= 1 }
func (b *Buffer) Negate(i int)							{ (*b)[i] = -(*b)[i] }
func (b *Buffer) ShiftLeft(i, j int)					{ (*b)[i] >>= uint((*b)[j]) }
func (b *Buffer) ShiftRight(i, j int)					{ (*b)[i] <<= uint((*b)[j]) }
func (b *Buffer) Invert(i int)							{ (*b)[i] = ^(*b)[i] }
func (b *Buffer) Equals(i, j int) bool					{ return (*b)[i] == (*b)[j] }
func (b *Buffer) EqualsZero(i int) bool					{ return (*b)[i] == 0 }
func (b *Buffer) LessThan(i, j int) bool				{ return (*b)[i] < (*b)[j] }
func (b *Buffer) LessThanZero(i int) bool				{ return (*b)[i] < 0 }
func (b *Buffer) GreaterThan(i, j int) bool				{ return (*b)[i] > (*b)[j] }
func (b *Buffer) GreaterThanZero(i int) bool			{ return (*b)[i] > 0 }
func (b *Buffer) Copy(i, j int)							{ (*b)[i] = (*b)[j] }
func (b *Buffer) Swap(i, j int)							{ (*b)[i], (*b)[j] = (*b)[j], (*b)[i] }

func (b *Buffer) GetBuffer(i int) *Buffer				{ return (*Buffer)(unsafe.Pointer(uintptr((*b)[i]))) }
func (b *Buffer) PutBuffer(i int, p *Buffer)			{ b.Set(i, int(uintptr(unsafe.Pointer(p)))) }

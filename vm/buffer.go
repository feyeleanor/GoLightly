//	TODO:	Storing and retrieving of pointers
//	TODO:	Shouldn't support Extend or Expand
//	TODO:	Interface?

package vm

import . "container/vector"
import "unsafe"

func max(x, y int) int									{ if x > y { return x}; return y }
func min(x, y int) int									{ if x < y { return x}; return y }

type Buffer struct {
	IntVector
}

func (b *Buffer) Init(length int)						{ b.IntVector.Resize(length, length) }
func (b *Buffer) Clear()								{ for i, _ := range b.IntVector { b.Set(i, 0) } }
func (b *Buffer) Cap() int								{ return b.Len() }
func (b *Buffer) Slice(i, j int) *Buffer				{ return &Buffer{*b.IntVector.Slice(i, j)} }
func (b *Buffer) Clone() *Buffer						{ return b.Slice(0, b.Len()) }

func (b *Buffer) Identical(o *Buffer) bool {
	match := b.Len() == o.Len()
	if match { for i, e := range o.IntVector { match = match && b.At(i) == e } }
	return match
}

func (b *Buffer) Replace(o *Buffer) {
	a := b.IntVector
	copy(a, o.Data())
}

func (b *Buffer) Add(i, j int)							{ a := b.IntVector; a[i] += a[j] }
func (b *Buffer) Subtract(i, j int)						{ a := b.IntVector; a[i] -= a[j] }
func (b *Buffer) Multiply(i, j int)						{ a := b.IntVector; a[i] *= a[j] }
func (b *Buffer) Divide(i, j int)						{ a := b.IntVector; a[i] /= a[j] }
func (b *Buffer) And(i, j int)							{ a := b.IntVector; a[i] &= a[j] }
func (b *Buffer) Or(i, j int)							{ a := b.IntVector; a[i] |= a[j] }
func (b *Buffer) Xor(i, j int)							{ a := b.IntVector; a[i] ^= a[j] }
func (b *Buffer) Increment(i int)						{ a := b.IntVector; a[i] += 1 }
func (b *Buffer) Decrement(i int)						{ a := b.IntVector; a[i] -= 1 }
func (b *Buffer) Negate(i int)							{ a := b.IntVector; a[i] = -a[i] }
func (b *Buffer) ShiftLeft(i, j int)					{ a := b.IntVector; a[i] >>= uint(a[j]) }
func (b *Buffer) ShiftRight(i, j int)					{ a := b.IntVector; a[i] <<= uint(a[j]) }
func (b *Buffer) Invert(i int)							{ a := b.IntVector; a[i] = ^a[i] }
func (b *Buffer) Equals(i, j int) bool					{ a := b.IntVector; return a[i] == a[j] }
func (b *Buffer) EqualsZero(i int) bool					{ a := b.IntVector; return a[i] == 0 }
func (b *Buffer) LessThan(i, j int) bool				{ a := b.IntVector; return a[i] < a[j] }
func (b *Buffer) LessThanZero(i int) bool				{ a := b.IntVector; return a[i] < 0 }
func (b *Buffer) GreaterThan(i, j int) bool				{ a := b.IntVector; return a[i] > a[j] }
func (b *Buffer) GreaterThanZero(i int) bool			{ a := b.IntVector; return a[i] > 0 }
func (b *Buffer) Copy(i, j int)							{ a := b.IntVector; a[i] = a[j] }

func (b *Buffer) GetBuffer(i int) *Buffer				{ return (*Buffer)(unsafe.Pointer(uintptr(b.At(i)))) }
func (b *Buffer) PutBuffer(i int, p *Buffer)			{ b.Set(i, int(uintptr(unsafe.Pointer(p)))) }

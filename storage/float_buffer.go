package storage

import "math"
//import "reflect"
//import "unsafe"

type FloatBuffer []float

func (b FloatBuffer) Len() int							{ return len(b) }
func (b FloatBuffer) Cap() int							{ return cap(b) }

func (b FloatBuffer) Clone() FloatBuffer {
	s := make(FloatBuffer, len(b))
	copy(s, b)
	return s
}

func (b *FloatBuffer) Resize(length int) {
	if length > cap(*b) {
		x := *b
		*b = make(FloatBuffer, length)
		copy(*b, x)
	} else {
		*b = (*b)[:length]
	}
}

func (b *FloatBuffer) Extend(count int) {
	b.Resize(len(*b) + count)
}

func (b *FloatBuffer) Shrink(count int) {
	b.Resize(len(*b) - count)
}

func (b FloatBuffer) Replicate(count int) FloatBuffer {
	l := len(b)
	s := make(FloatBuffer, l * count)
	for i, offset := 0, 0; i < count; i++ {
		copy(s[offset:], b)
		offset += l
	}
	return s
}

func (b FloatBuffer) Collect(f func(i int, x float) float) FloatBuffer {
	n := make(FloatBuffer, len(b))
	for i, x := range b {
		n[i] = f(i, x)
	}
	return n
}

func (b FloatBuffer) Copy(i, j int)						{ b[i] = b[j] }
func (b FloatBuffer) Swap(i, j int)						{ b[i], b[j] = b[j], b[i] }
func (b FloatBuffer) Clear(i, j int)					{ for ; i < j; i++ { b[i] = float(0) } }

func (b FloatBuffer) Negate(i int)						{ b[i] = -b[i] }
func (b FloatBuffer) Increment(i int)					{ b[i] += 1 }
func (b FloatBuffer) Decrement(i int)					{ b[i] -= 1 }

func (b FloatBuffer) Add(i, j int)						{ b[i] += b[j] }
func (b FloatBuffer) Subtract(i, j int)					{ b[i] -= b[j] }
func (b FloatBuffer) Multiply(i, j int)					{ b[i] *= b[j] }
func (b FloatBuffer) Divide(i, j int)					{ b[i] /= b[j] }
func (b FloatBuffer) Remainder(i, j int)				{ b.Clear(i, 1) }

func (b FloatBuffer) Equals(i, j int, t float) bool		{ return math.Fabs(float64(b[i] - b[j])) < float64(t) }
func (b FloatBuffer) EqualsZero(i int, t float) bool	{ return b[i] < t }
func (b FloatBuffer) LessThan(i, j int) bool			{ return b[i] < b[j] }
func (b FloatBuffer) LessThanZero(i int) bool			{ return b[i] < float(0) }
func (b FloatBuffer) GreaterThan(i, j int) bool			{ return b[i] > b[j] }
func (b FloatBuffer) GreaterThanZero(i int) bool		{ return b[i] > float(0) }

func (b FloatBuffer) Feed(c chan<- float, f func(i int, x float) float) {
	d := b.Clone()
	go func() {
		for i, v := range d { c <- f(i, v) }
		close(c)
	}()
}

func (b FloatBuffer) Pipe(f func(i int, x float) float) <-chan float {
	c := make(chan float)
	b.Feed(c, f)
	return c
}

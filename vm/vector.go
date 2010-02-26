//	TODO:	Storing and retrieving of pointers
//	TODO:	Shouldn't support Extend or Expand
//	TODO:	Interface?

package vm

import "unsafe"
import "math"

type Vector struct {
	Buffer
}

func (s *Vector) Slice(i, j int) *Vector				{ return &Vector{*s.Buffer.Slice(i, j)} }
func (s *Vector) Clone() *Vector						{ return s.Slice(0, s.Len()) }
func (s *Vector) Identical(o *Vector) bool				{ return s.Buffer.Identical(&o.Buffer) }
func (s *Vector) FIdentical(o *Vector, t float) bool	{ return s.Buffer.FIdentical(&o.Buffer, t) }
func (s *Vector) Replace(o *Vector)						{ s.Buffer.Replace(&o.Buffer) }

func (s *Vector) Add(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] += o.Buffer[i] }
}

func (s *Vector) Subtract(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] -= o.Buffer[i] }
}

func (s *Vector) Multiply(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] *= o.Buffer[i] }
}

func (s *Vector) Divide(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] /= o.Buffer[i] }
}

func (s *Vector) FAdd(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.Buffer))
	for i := 0; i < limit; i++ { sfb[i + offset] += ofb[i] }
}

func (s *Vector) FSubtract(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.Buffer))
	for i := 0; i < limit; i++ { sfb[i + offset] -= ofb[i] }
}

func (s *Vector) FMultiply(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.Buffer))
	for i := 0; i < limit; i++ { sfb[i + offset] *= ofb[i] }
}

func (s *Vector) FDivide(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.Buffer))
	for i := 0; i < limit; i++ { sfb[i + offset] /= ofb[i] }
}

func (s *Vector) And(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] &= o.Buffer[i] }
}

func (s *Vector) Or(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] |= o.Buffer[i] }
}

func (s *Vector) Xor(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] ^= o.Buffer[i] }
}

func (s *Vector) Clear(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] = 0 }
}

func (s *Vector) ClearAll() {
	for i := range s.Buffer { s.Buffer[i] = 0 }
}

func (s *Vector) Increment(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] += 1 }
}

func (s *Vector) IncrementAll() {
	for i := range s.Buffer { s.Buffer[i] += 1 }
}

func (s *Vector) Decrement(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] -= 1 }
}

func (s *Vector) DecrementAll() {
	for i := range s.Buffer { s.Buffer[i] -= 1 }
}

func (s *Vector) Negate(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] = -s.Buffer[i] }
}

func (s *Vector) NegateAll() {
	for i, e := range s.Buffer { s.Buffer[i] = -e }
}

func (s *Vector) FNegate(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	for i := offset; i < limit; i++ { sfb[i] = -sfb[i] }
}

func (s *Vector) FNegateAll() {
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	for i, _ := range s.Buffer { sfb[i] = -sfb[i] }
}

func (s *Vector) ShiftLeft(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] >>= uint(o.Buffer[i]) }
}

func (s *Vector) ShiftRight(offset int, o *Vector) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] <<= uint(o.Buffer[i]) }
}

func (s *Vector) Invert(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] = ^s.Buffer[i] }
}

func (s *Vector) Equals(offset int, o *Vector) bool {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	t := true
	for i := 0; i < limit; i++ { t = t && s.Buffer[i + offset] == o.Buffer[i] }
	return t
}

func (s *Vector) EqualsZero(offset, count int) bool {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	t := true
	for i := offset; t && i < limit; i++ { t = t && s.Buffer[i] == 0 }
	return t
}

func (s *Vector) LessThan(offset int, o *Vector) bool {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	t := true
	for i := 0; i < limit; i++ { t = t && s.Buffer[i + offset] < o.Buffer[i] }
	return t
}

func (s *Vector) LessThanZero(offset, count int) bool {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	t := true
	for i := offset; t && i < limit; i++ { t = t && s.Buffer[i] < 0 }
	return t
}

func (s *Vector) GreaterThan(offset int, o *Vector) bool {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	t := true
	for i := 0; i < limit; i++ { t = t && s.Buffer[i + offset] > o.Buffer[i] }
	return t
}

func (s *Vector) GreaterThanZero(offset, count int) bool {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	t := true
	for i := offset; t && i < limit; i++ { t = t && s.Buffer[i] > 0 }
	return t
}

func (s *Vector) FEquals(offset int, o *Vector, tolerance float) bool {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	t := true
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.Buffer))
	for i := 0; i < limit; i++ { t = t && math.Fabs(float64(sfb[i + offset] - ofb[i])) < float64(tolerance) }
	return t
}

func (s *Vector) FEqualsZero(offset, count int, tolerance float) bool {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	t := true
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	for i := offset; t && i < limit; i++ { t = t && math.Fabs(float64(sfb[i])) < float64(tolerance) }
	return t
}

func (s *Vector) FLessThan(offset int, o *Vector) bool {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.Buffer))
	t := true
	for i := 0; i < limit; i++ { t = t && sfb[i + offset] < ofb[i] }
	return t
}

func (s *Vector) FLessThanZero(offset, count int) bool {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	t := true
	for i := offset; t && i < limit; i++ { t = t && sfb[i] < 0 }
	return t
}

func (s *Vector) FGreaterThan(offset int, o *Vector) bool {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	ofb := *(*[]float)(unsafe.Pointer(&o.Buffer))
	t := true
	for i := 0; i < limit; i++ { t = t && sfb[i + offset] > ofb[i] }
	return t
}

func (s *Vector) FGreaterThanZero(offset, count int) bool {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s.Buffer))
	t := true
	for i := offset; t && i < limit; i++ { t = t && sfb[i] > 0 }
	return t
}

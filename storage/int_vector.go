//	TODO:	Storing and retrieving of pointers
//	TODO:	Shouldn't support Extend or Expand
//	TODO:	Interface?

package storage

import "reflect"
import "unsafe"
import "math"

type IntVector []int

func (s IntVector) Clone() IntVector {
	n := make(IntVector, len(s))
	copy(n, s)
	return n
}
func (s IntVector) Identical(o IntVector) bool				{ return reflect.DeepEqual(s, o) }
func (s IntVector) Replace(o IntVector)						{ copy(s, o) }
func (s IntVector) Add(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] += v }
}

func (s IntVector) Subtract(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] -= v }
}

func (s IntVector) Multiply(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] *= v }
}

func (s IntVector) Divide(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] /= v }
}

func (s IntVector) FAdd(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	s = s[offset:limit]
	o = o[:limit]
	sfb := *(*[]float)(unsafe.Pointer(&s))
	ofb := *(*[]float)(unsafe.Pointer(&o))
	for i, v := range ofb { sfb[i] += v }
}

func (s IntVector) FSubtract(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	s = s[offset:limit]
	o = o[:limit]
	sfb := *(*[]float)(unsafe.Pointer(&s))
	ofb := *(*[]float)(unsafe.Pointer(&o))
	for i, v := range ofb { sfb[i] -= v }
}

func (s IntVector) FMultiply(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	s = s[offset:limit]
	o = o[:limit]
	sfb := *(*[]float)(unsafe.Pointer(&s))
	ofb := *(*[]float)(unsafe.Pointer(&o))
	for i, v := range ofb { sfb[i] *= v }
}

func (s IntVector) FDivide(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	s = s[offset:limit]
	o = o[:limit]
	sfb := *(*[]float)(unsafe.Pointer(&s))
	ofb := *(*[]float)(unsafe.Pointer(&o))
	for i, v := range ofb { sfb[i] /= v }
}

func (s IntVector) And(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] &= v }
}

func (s IntVector) Or(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] |= v }
}

func (s IntVector) Xor(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] ^= v }
}

func (s IntVector) Clear(offset, count int) {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	for i := range s[offset:limit] { s[i] = 0 }
}

func (s IntVector) ClearAll() {
	for i := range s { s[i] = 0 }
}

func (s IntVector) Increment(offset, count int) {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	for i := range s[offset:limit] { s[i] += 1 }
}

func (s IntVector) IncrementAll() {
	for i := range s { s[i] += 1 }
}

func (s IntVector) Decrement(offset, count int) {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	for i := range s[offset:limit] { s[i] -= 1 }
}

func (s IntVector) DecrementAll() {
	for i := range s { s[i] -= 1 }
}

func (s IntVector) Negate(offset, count int) {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	for i, v := range s[offset:limit] { s[i] = -v }
}

func (s IntVector) NegateAll() {
	for i, v := range s { s[i] = -v }
}

func (s IntVector) FNegate(offset, count int) {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s))
	for i, v := range sfb[offset:limit] { sfb[i] = -v }
}

func (s IntVector) FNegateAll() {
	sfb := *(*[]float)(unsafe.Pointer(&s))
	for i, v := range sfb { sfb[i] = -v }
}

func (s IntVector) ShiftLeft(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] <<= uint(v) }
}

func (s IntVector) ShiftRight(offset int, o IntVector) {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	for i, v := range o[:limit] { sb[i] >>= uint(v) }
}

func (s IntVector) Invert(offset, count int) {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	for i, v := range s[offset:limit] { s[i] = ^v }
}

func (s IntVector) Equals(offset int, o IntVector) bool {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	t := true
	for i, v := range o[:limit] {
		if t {
			t = t && sb[i] == v
		} else {
			break
		}
	}
	return t
}

func (s IntVector) EqualsZero(offset, count int) bool {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	t := true
	for _, v := range s[offset:limit] {
		if t {
			t = t && v == 0
		} else {
			break
		}
	}
	return t
}

func (s IntVector) LessThan(offset int, o IntVector) bool {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	t := true
	for i, v := range o[:limit] {
		if t {
			t = t && sb[i] < v
		} else {
			break
		}
	}
	return t
}

func (s IntVector) LessThanZero(offset, count int) bool {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	t := true
	for _, v := range s[offset:limit] {
		if t {
			t = t && v < 0
		} else {
			break
		}
	}
	return t
}

func (s IntVector) GreaterThan(offset int, o IntVector) bool {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sb := s[offset:limit]
	t := true
	for i, v := range o[:limit] {
		if t {
			t = t && sb[i] > v
		} else {
			break
		}
	}
	return t
}

func (s IntVector) GreaterThanZero(offset, count int) bool {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	t := true
	for _, v := range s[offset:limit] {
		if t {
			t = t && v > 0
		} else {
			break
		}
	}
	return t
}

func (s IntVector) FEquals(offset int, o IntVector, tolerance float) bool {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	t := true
	tol := math.Fabs(float64(tolerance))
	sfb := *(*[]float)(unsafe.Pointer(&s))
	sfb = sfb[offset:limit]
	ofb := *(*[]float)(unsafe.Pointer(&o))
	for i, v := range ofb[:limit] {
		if t {
			t = t && math.Fabs(float64(sfb[i] - v)) < tol
		} else {
			break
		}
	}
	return t
}

func (s IntVector) FEqualsZero(offset, count int, tolerance float) bool {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	t := true
	tol := math.Fabs(float64(tolerance))
	sfb := *(*[]float)(unsafe.Pointer(&s))
	for _, v := range sfb[offset:limit] {
		if t {
			t = t && math.Fabs(float64(v)) < tol
		} else {
			break
		}
	}
	return t
}

func (s IntVector) FLessThan(offset int, o IntVector) bool {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s))
	sfb = sfb[offset:limit]
	ofb := *(*[]float)(unsafe.Pointer(&o))
	t := true
	for i, v := range ofb[:offset] {
		if t {
			t = t && sfb[i] < v
		} else {
			break
		}
	}
	return t
}

func (s IntVector) FLessThanZero(offset, count int) bool {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s))
	t := true
	for _, v := range sfb[offset:limit] {
		if t {
			t = t && v < 0
		} else {
			break
		}
	}
	return t
}

func (s IntVector) FGreaterThan(offset int, o IntVector) bool {
	var limit int
	if len(s) < (len(o) + offset) {
		limit = len(s)
	} else {
		limit = len(o) + offset
	}
	sfb := *(*[]float)(unsafe.Pointer(&s))
	sfb = sfb[offset:limit]
	ofb := *(*[]float)(unsafe.Pointer(&o))
	t := true
	for i, v := range ofb[:limit] {
		if t{
			t = t && sfb[i] > v
		} else {
			break
		}
	}
	return t
}

func (s IntVector) FGreaterThanZero(offset, count int) bool {
	var limit int
	if len(s) < (offset + count) {
		limit = len(s)
	} else {
		limit = offset + count
	}
	sfb := *(*[]float)(unsafe.Pointer(&s))
	t := true
	for _, v := range sfb[offset:limit] {
		if t {
			t = t && v > 0
		} else {
			break
		}
	}
	return t
}
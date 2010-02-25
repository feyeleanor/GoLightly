//	TODO:	Storing and retrieving of pointers
//	TODO:	Shouldn't support Extend or Expand
//	TODO:	Interface?

package vm

import "unsafe"
//import "os"
//import "fmt"

type Stream struct {
	Buffer
}

func (s *Stream) Slice(i, j int) *Stream				{ return &Stream{*s.Buffer.Slice(i, j)} }
func (s *Stream) Clone() *Stream						{ return s.Slice(0, s.Len()) }
func (s *Stream) Identical(o *Stream) bool				{ return s.Buffer.Identical(&o.Buffer) }
func (s *Stream) Replace(o *Stream)						{ s.Buffer.Replace(&o.Buffer) }

func (s *Stream) Add(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] += o.Buffer[i] }
}

func (s *Stream) Subtract(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] -= o.Buffer[i] }
}

func (s *Stream) Multiply(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] *= o.Buffer[i] }
}

func (s *Stream) Divide(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] /= o.Buffer[i] }
}

func (s *Stream) FAdd(offset int, o *Stream) {
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

func (s *Stream) FSubtract(offset int, o *Stream) {
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

func (s *Stream) FMultiply(offset int, o *Stream) {
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

func (s *Stream) FDivide(offset int, o *Stream) {
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

func (s *Stream) And(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] &= o.Buffer[i] }
}

func (s *Stream) Or(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] |= o.Buffer[i] }
}

func (s *Stream) Xor(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] ^= o.Buffer[i] }
}

func (s *Stream) Clear(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] = 0 }
}

func (s *Stream) ClearAll() {
	for i := range s.Buffer { s.Buffer[i] = 0 }
}

func (s *Stream) Increment(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] += 1 }
}

func (s *Stream) IncrementAll() {
	for i := range s.Buffer { s.Buffer[i] += 1 }
}

func (s *Stream) Decrement(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] -= 1 }
}

func (s *Stream) DecrementAll() {
	for i := range s.Buffer { s.Buffer[i] -= 1 }
}

func (s *Stream) Negate(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] = -s.Buffer[i] }
}

func (s *Stream) NegateAll() {
	for i, e := range s.Buffer { s.Buffer[i] = -e }
}

func (s *Stream) ShiftLeft(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] >>= uint(o.Buffer[i]) }
}

func (s *Stream) ShiftRight(offset int, o *Stream) {
	var limit int
	if len(s.Buffer) < (len(o.Buffer) + offset) {
		limit = len(s.Buffer)
	} else {
		limit = len(o.Buffer) + offset
	}
	for i := 0; i < limit; i++ { s.Buffer[i + offset] <<= uint(o.Buffer[i]) }
}

func (s *Stream) Invert(offset, count int) {
	var limit int
	if len(s.Buffer) < (offset + count) {
		limit = len(s.Buffer)
	} else {
		limit = offset + count
	}
	for i := offset; i < limit; i++ { s.Buffer[i] = ^s.Buffer[i] }
}

func (s *Stream) Equals(offset int, o *Stream) bool {
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

func (s *Stream) EqualsZero(offset, count int) bool {
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

func (s *Stream) LessThan(offset int, o *Stream) bool {
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

func (s *Stream) LessThanZero(offset, count int) bool {
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

func (s *Stream) GreaterThan(offset int, o *Stream) bool {
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

func (s *Stream) GreaterThanZero(offset, count int) bool {
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

func (b *Stream) Copy(o *Stream) {
}

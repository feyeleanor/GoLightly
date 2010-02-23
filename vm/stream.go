//	TODO:	Storing and retrieving of pointers
//	TODO:	Shouldn't support Extend or Expand
//	TODO:	Interface?

package vm

//import . "container/vector"
//import "unsafe"

type Stream struct {
	Buffer
}

func (s *Stream) Init(length int)						{ s.Buffer.Init(length) }
func (s *Stream) Clear()								{ s.Buffer.Clear() }
func (s *Stream) Slice(i, j int) *Stream				{ return &Stream{*s.Buffer.Slice(i, j)} }
func (s *Stream) Clone() *Stream						{ return s.Slice(0, s.Len()) }

func (s *Stream) Identical(o *Stream) bool {
	return s.Buffer.Identical(&o.Buffer)
}

func (s *Stream) Replace(o *Stream) {
	s.Buffer.Replace(&o.Buffer)
}

func (s *Stream) Add(o *Stream) {
	x := s.Buffer.IntVector
	y := o.Buffer.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] += y[i] }
}

func (s *Stream) Subtract(o *Stream) {
	x := s.Buffer.IntVector
	y := o.Buffer.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] -= y[i] }
}

func (s *Stream) Multiply(o *Stream) {
	x := s.Buffer.IntVector
	y := o.Buffer.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] *= y[i] }
}

func (s *Stream) Divide(o *Stream) {
	x := s.Buffer.IntVector
	y := o.Buffer.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] /= y[i] }
}

func (s *Stream) And(o *Stream) {
	x := s.Buffer.IntVector
	y := o.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] &= y[i] }
}

func (s *Stream) Or(o *Stream) {
	x := s.Buffer.IntVector
	y := o.Buffer.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] |= y[i] }
}

func (s *Stream) Xor(o *Stream) {
	x := s.Buffer.IntVector
	y := o.Buffer.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] ^= y[i] }
}

func (s *Stream) Increment() {
	x := s.Buffer.IntVector
	for i := 0; i < len(x); i++ { x[i] += 1 }
}

func (s *Stream) Decrement() {
	x := s.Buffer.IntVector
	for i := 0; i < len(x); i++ { x[i] -= 1 }
}

func (s *Stream) Negate() {
	x := s.Buffer.IntVector
	for i := 0; i < len(x); i++ { x[i] = -x[i] }
}

func (s *Stream) ShiftLeft(o *Stream) {
	x := s.Buffer.IntVector
	y := o.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] >>= uint(y[i]) }
}

func (s *Stream) ShiftRight(o *Stream) {
	x := s.Buffer.IntVector
	y := o.IntVector
	r := min(len(x), len(y))
	for i := 0; i < r; i++ { x[i] <<= uint(y[i]) }
}

func (s *Stream) Invert() {
	x := s.Buffer.IntVector
	for i := 0; i < len(x); i++ { x[i] = ^x[i] }
}

func (s *Stream) Equals(o *Stream) bool {
	x := s.Buffer.IntVector
	y := o.IntVector
	r := min(len(x), len(y))
	t := true
	for i := 0; t && i < r; i++ { t = t && x[i] == y[i] }
	return t
}

func (s *Stream) EqualsZero() bool {
	t := true
	x := s.Buffer.IntVector
	for i := 0; t && i < len(x); i++ { t = t && x[i] == 0 }
	return t
}

func (s *Stream) LessThan(o *Stream) bool {
	x := s.Buffer.IntVector
	y := o.IntVector
	r := min(len(x), len(y))
	t := true
	for i := 0; t && i < r; i++ { t = t && x[i] < y[i] }
	return t
}

func (s *Stream) LessThanZero() bool {
	t := true
	x := s.Buffer.IntVector
	for i := 0; t && i < len(x); i++ { t = t && x[i] < 0 }
	return t
}

func (s *Stream) GreaterThan(o *Stream) bool {
	x := s.Buffer.IntVector
	y := o.IntVector
	r := min(len(x), len(y))
	t := true
	for i := 0; t && i < r; i++ { t = t && x[i] > y[i] }
	return t
}

func (s *Stream) GreaterThanZero() bool {
	t := true
	x := s.Buffer.IntVector
	for i := 0; t && i < len(x); i++ { t = t && x[i] > 0 }
	return t
}

func (b *Stream) Copy(o *Stream) {
}

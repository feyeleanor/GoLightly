package testLightly

import "fmt"
import "path"
import "strings"
import "reflect"
import "runtime"
import "testing"

func formatMessage(header interface{}, statements... interface{}) string {
	return fmt.Sprintf("%v: %v", header, fmt.Sprint(statements...))
}

func describeMismatch(target, value interface{}) string {
	return fmt.Sprintf("%v != %v", target, value)
}

type SimpleClosure			func()
type BlockingFunction		func(done chan bool)
type TestWrapper			func(T *Test)

type Test struct {
	T			*testing.T
	title		string
	location	string
	captioned	bool
	tracing		bool
	stack_trace	[]uintptr
}

func NewTest(t *testing.T) *Test {
	return &Test{T:t, tracing: true}
}

func (t *Test) recordLocation() *Test {
	runtime.Callers(1, t.stack_trace)
	my_path := ""
	for i, u := range t.stack_trace {
		if f := runtime.FuncForPC(u); f != nil {
			file, line := f.FileLine(u)
			if i == 0 {
				my_path = file
			} else if my_path != file {
				t.location = fmt.Sprintf("%v:%v", path.Base(file), line)
				break
			}
		} else {
			t.location = "(unknown)"
			continue
		}
	}
	return t
}

func (t *Test) LogHeader() {
	t.recordLocation()
	if !t.captioned {
		boxing := fmt.Sprintf("+--%v--+", strings.Repeat("-", len(t.title)))
		t.T.Log(boxing)
		t.T.Log("| ", t.title, " |")
		t.T.Log(boxing)
		t.captioned = true
	}
}

func (t *Test) Log(comments... interface{}) {
	t.LogHeader()
	t.T.Log("  ", fmt.Sprint(comments...))
}

func (t *Test) Error(comments... interface{}) {
	t.LogHeader()
	t.T.Error("  ", t.location, "-->", fmt.Sprint(comments...))
}

func (t *Test) Comment(comments... interface{}) *Test {
	t.Log(comments...)
	return t
}

func (t *Test) Run(title string, f TestWrapper) *Test {
	t.location = ""
	t.title = title
	t.captioned = false
	t.stack_trace = make([]uintptr, 100)
	f(t)
	return t
}

func (t *Test) ToDo(s string) *Test {
	t.Comment("[TO DO] ", s)
	return t
}

func (t *Test) Untested(s string) *Test {
	t.ToDo(fmt.Sprintf("Write tests for %v()", s))
	return t
}

func (t *Test) Unbenchmarked(s string) *Test {
	t.ToDo(fmt.Sprintf("Write benchmarks for %v()", s))
	return t
}

func (t *Test) Unimplemented(s string) *Test {
	t.ToDo(fmt.Sprintf("Implement %v()", s))
	return t
}

func (t *Test) Confirm(values... bool) *Test {
	for i, v := range values {
		if !v { t.Error(i + 1, " > ", describeMismatch(true, false)) }
	}
	return t
}

func (t *Test) Refute(values... bool) *Test {
	for i, v := range values {
		if v { t.Error(i + 1, " > ", describeMismatch(false, true)) }
	}
	return t
}

func (t *Test) EachPermutation(values []interface{}, f func(id string, x, y interface{})) {
	for i, x := range values {
		for j, y := range values {
			if j > i {
				f(fmt.Sprintf("[%v, %v] ", i, j), x, y)
			}
		}
	}
}

func (t *Test) Identical(values... interface{}) *Test {
	t.EachPermutation(values, func(id string, x, y interface{}) {
		if !reflect.DeepEqual(x, y) {
			t.Error(id, describeMismatch(x, y))
		}
	})
	return t
}

func (t *Test) Different(values... interface{}) *Test {
	t.EachPermutation(values, func(id string, x, y interface{}) {
		if reflect.DeepEqual(x, y) {
			t.Error(id, fmt.Sprintf("%v == %v", x, y))
		}
	})
	return t
}

func (t *Test) printStackTrace() *Test {
	if t.tracing {
		
	}
	return t
}

func (t *Test) Erroneous(f SimpleClosure) *Test {
	defer func() {
		if r := recover(); r == nil {
			t.Error("should cause a panic")
			t.printStackTrace()
		}
	}()
	f()
	return t
}

func (t *Test) Safe(f SimpleClosure) *Test {
	defer func() {
		if r := recover(); r != nil {
			t.Error("shouldn't cause a panic")
			t.printStackTrace()
		}
	}()
	f()
	return t
}

func (t *Test) WaitFor(f BlockingFunction) *Test {
	done := make(chan bool)
	go f(done)
	<-done
	return t
}

package testLightly

import "fmt"

type Reduction		func(y, x interface{}) interface{}
type ValueSet		[]interface{}

type TestTable struct {
	XValues			ValueSet
	YValues			ValueSet
	compatibility	[]ValueSet
	test			Reduction
}

func NewTestTable(test Reduction) *TestTable {
	t := new(TestTable)
	t.test = test
	return t
}

func (t *TestTable) X(values... interface{}) *TestTable {
	t.XValues = append(t.XValues, values...)
	return t
}

func (t *TestTable) Y(value interface{}, compatibility... interface{}) *TestTable {
	t.YValues = append(t.YValues, value)
	t.compatibility = append(t.compatibility, append(ValueSet{}, compatibility...))
	return t
}

func (t *TestTable) Assess(T *Test) *TestTable {
	for row, x := range t.YValues {
		for column, y := range t.XValues {
			expectation := t.compatibility[row][column]
			if result := t.Apply(x, y); result != expectation {
				T.Error(fmt.Sprintf("[%v, %v]", row, column), fmt.Sprintf("-> expected %v got %v", expectation, result))
			}
		}
	}
	return t
}

func (t *TestTable) Apply(x, y interface{}) (i interface{}) {
	defer func() { if recover() != nil { i = nil } }()
	return t.test(x, y)
}
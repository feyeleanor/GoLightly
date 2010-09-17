package testLightly

import "fmt"

type Reduction				func(y, x interface{}) interface{}

type ResultSet map[interface{}] interface{}

type TestTable struct {
	values			[]interface{}
	compatibility	map[interface{}] ResultSet
	test			Reduction
}

func NewTestTable(test Reduction, values... interface{}) *TestTable {
	t := new(TestTable)
	t.test = test
	t.X(values)
	return t
}

func (t *TestTable) X(values... interface{}) *TestTable {
	t.values = values
	t.compatibility = make(map[interface{}] ResultSet)
	for _, v := range values {
		t.compatibility[v] = make(ResultSet)
	}
	return t
}
func (t *TestTable) Y(value interface{}, compatibility... interface{}) *TestTable {
	if len(compatibility) == len(t.values) {
		for i, v := range t.values {
			if column, ok := t.compatibility[v]; ok {
				column[value] = compatibility[i]
			} else {
				panic(t)
			}
		}
		return t
	}
	panic(t)
}
func (t *TestTable) Assess(T *Test) *TestTable {
	for column, tests := range t.compatibility {
		for row, expectation := range tests {
			if result := t.Apply(row, column); result != expectation {
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

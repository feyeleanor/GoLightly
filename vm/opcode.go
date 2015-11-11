package vm

import (
	"fmt"
	"reflect"
)

type OpCode struct {
	Code			int
	Movement	int
	Data			interface{}
}

func (o OpCode) Similar(p OpCode) bool {
	return o.Code == p.Code && o.Movement == p.Movement && reflect.TypeOf(o.Data) == reflect.TypeOf(p.Data)
}

func (o OpCode) Identical(p OpCode) bool {
	return reflect.DeepEqual(o, p)
}

func (o *OpCode) Replace(p *OpCode) {
	o.Code = p.Code
	o.Movement = p.Movement
	o.Data = p.Data
}

func (o *OpCode) String() string {
	return fmt.Sprintf("%v: %v", o.Code, o.Data)
}
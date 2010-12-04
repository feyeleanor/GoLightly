package filters

import . "reflect"

type Enumerable interface {
	Each(f interface{})
}

func unpackValueArray(v []Value) []interface{} {
	a := make([]interface{}, len(v))
	for i, x := range v { a[i] = x }
	return a
}

func Each(i, f interface{}) {
	switch i := i.(type) {
	case Enumerable:
		i.Each(f)
	case []interface{}:
		if f, ok := f.(func(interface{})); ok {
			for _, v := range i { f(v) }
		}
	default:
		switch f := f.(type) {
		case func(interface{}):
			switch i := NewValue(i).(type) {
			case *ChanValue:
				for !i.Closed() { f(i.Recv().Interface()) }
			case *StringValue:
				for c := range i.Get() { f(c) }
			case *SliceValue:
				for j := 0; j < i.Len(); j++ { f(i.Elem(j).Interface()) }
			case *MapValue:
				for _, k := range i.Keys() { f(i.Elem(k).Interface()) }
			case *FuncValue:
				for {
					if v := i.Call([]Value{}); len(v) != 0 {
						f(unpackValueArray(v))
					} else {
						break
					}
				}
			default:
				f(i.Interface())
			}
		default:
			fullyReflectedEach(NewValue(i), NewValue(f))
		}
	}
}

func fullyReflectedEach(i, f Value) {
	switch f := f.(type) {
	case *FuncValue:
		if t := f.Type().(*FuncType); t.NumIn() != 0 {
			switch i := i.(type) {
			case *StringValue:
				if t == Typeof(int(0)) {
					for c := range i.Get() {
						f.Call([]Value{NewValue(c)})
					}
				}
			case *ChanValue:
				for !i.Closed() {
					f.Call([]Value{i.Recv()})
				}
			case *SliceValue:
				if i.Type().(*SliceType).Elem() == t.In(0) {
					for j := 0; j < i.Len(); j++ {
						f.Call([]Value{i.Elem(j)})
					}
				}
			case *MapValue:
				for _, k := range i.Keys() {
					f.Call([]Value{i.Elem(k)})
				}
			case *FuncValue:
				for {
					if i.Type().(*FuncType).Out(0) == t.In(0) {
						if v := i.Call([]Value{}); len(v) != 0 {
							f.Call(v)
						} else {
							break
						}
					}
				}
			}
		}
	}
}
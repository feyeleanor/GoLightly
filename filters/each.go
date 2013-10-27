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
			switch i := ValueOf(i); i.Kind() {
			case Chan:
				for v, ok := i.Recv(); ok; v, ok = i.Recv() { f(v.Interface()) }
			case String, Slice:
				for j := 0; j < i.Len(); j++ { f(i.Index(j).Interface()) }
			case Map:
				for _, k := range i.MapKeys() { f(i.MapIndex(k).Interface()) }
			case Func:
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
			fullyReflectedEach(ValueOf(i), ValueOf(f))
		}
	}
}

func fullyReflectedEach(i, f Value) {
	switch f.Kind() {
	case Func:
		if t := TypeOf(f); t.NumIn() != 0 {
			switch i.Kind() {
			case Chan:
				for v, ok := i.Recv(); ok; v, ok = i.Recv() {
					f.Call([]Value{v})
				}
			case String:
				for j := 0; j < i.Len(); j++ {
					f.Call([]Value{i.Index(j)})
				}
			case Slice:
				for j := 0; j < i.Len(); j++ {
					f.Call([]Value{i.Index(j)})
				}
			case Map:
				for _, k := range i.MapKeys() {
					f.Call([]Value{i.MapIndex(k)})
				}
			case Func:
				for v := i.Call([]Value{}); len(v) != 0; v = i.Call([]Value{}) {
					f.Call(v)
				}
			}
		}
	}
}
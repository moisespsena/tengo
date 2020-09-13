package tengo

import (
	"fmt"
	"reflect"

	"github.com/d5/tengo/v2/token"
)

type Struct struct {
	ObjectImpl
	Fields, Funcs *Map
	Methods       map[string]*StructMethod
	Name          string
}

func NewStruct(fields *Map, funcs *Map) (*Struct, error) {
	if fields == nil {
		fields = &Map{Value: map[string]Object{}}
	}
	if funcs == nil {
		funcs = &Map{Value: map[string]Object{}}
	}
	s := &Struct{
		Fields:  fields,
		Funcs:   funcs,
		Methods: map[string]*StructMethod{},
	}
	for key, f := range funcs.Value {
		if !f.CanCall() {
			return nil, ErrInvalidMapIndexValueType{
				MapName:   "funcs",
				IndexName: key,
				Expected:  "callable",
				Found:     "not callable",
			}
		}
		s.Methods[key] = &StructMethod{
			Func:   f,
			Name:   key,
			Struct: s,
		}
	}
	return s, nil
}

func (s *Struct) TypeName() string {
	return "struct"
}

func (s *Struct) String() string {
	return "<struct of " + s.Fields.String() + ">"
}

func (s *Struct) IsFalsy() bool {
	return false
}

func (s *Struct) Equals(another Object) bool {
	if sm, ok := another.(*Struct); ok {
		return sm == s
	}
	return false
}

func (s *Struct) Copy() Object {
	return s.copy()
}

func (s *Struct) copy() *Struct {
	copy, _ := NewStruct(s.Fields, s.Funcs)
	return copy
}

func (s *Struct) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := ToString(index)
	if !ok {
		err = ErrInvalidIndexType
		return
	}

	if strIdx == "__map__" {
		res = &Map{Value: map[string]Object{
			"fields": s.Fields,
			"funcs":  s.Funcs,
		}}
		return
	}

	if res, ok = s.Methods[strIdx]; !ok {
		res, ok = s.Fields.Value[strIdx]
		if !ok {
			res = UndefinedValue
		}
	}
	return
}

func (s *Struct) IndexSet(index, value Object) error {
	strIdx, ok := ToString(index)
	if !ok {
		return ErrInvalidIndexType
	}
	s.Fields.Value[strIdx] = value
	return nil
}

func (s *Struct) Iterate() Iterator {
	return s.Fields.Iterate()
}

func (s *Struct) CanIterate() bool {
	return true
}

func (s *Struct) Call(args ...Object) (ret Object, err error) {
	copy := s.copy()
	copy.Fields = copy.Fields.Copy().(*Map)
	if m := copy.Methods["__constructor"]; m != nil {
		if _, err = m.Call(args...); err != nil {
			return
		}
	}
	return &StructInstance{copy}, nil
}

func (s *Struct) CanCall() bool {
	return true
}

func (s *Struct) CanCallContext() bool {
	return true
}

type StructInstance struct {
	*Struct
}

func (i *StructInstance) Call(args ...Object) (ret Object, err error) {
	return i.Methods["__call"].Call(args...)
}

func (i *StructInstance) CanCall() bool {
	return i.Methods["__call"] != nil
}

func (i *StructInstance) CanCallContext() bool {
	return true
}
func (StructInstance) TypeName() string {
	return "struct-instance"
}

func (i *StructInstance) String() string {
	return "<struct-instance of " + i.Fields.String() + ">"
}

type StructMethod struct {
	ObjectImpl
	Func   Object
	Name   string
	Struct *Struct
}

func (m *StructMethod) TypeName() string {
	return "struct_method"
}

func (m *StructMethod) String() string {
	return "<struct_method '" + m.Name + "'>"
}

func (m *StructMethod) BinaryOp(op token.Token, rhs Object) (Object, error) {
	return nil, ErrInvalidOperator
}

func (m *StructMethod) IsFalsy() bool {
	return false
}

func (m *StructMethod) Equals(another Object) bool {
	if anotherM, ok := another.(*StructMethod); ok {
		return anotherM.Name == m.Name && anotherM.Struct == m.Struct
	}
	return false
}

func (m StructMethod) Copy() Object {
	return &m
}

func (m *StructMethod) Call(args ...Object) (ret Object, err error) {
	return Call(args[0].(*Context), m.Func, append([]Object{m.Struct}, args[1:]...)...)
}

func (m *StructMethod) CanCall() bool {
	return true
}

func (m *StructMethod) CanCallContext() bool {
	return true
}

type ReflectedStruct struct {
	ObjectImpl
	Typ        reflect.Type
	Fields     map[string][]int
	Methods    map[string]int
	Construtor int
}

var structTypes = map[reflect.Type]*ReflectedStruct{}

func NewReflectedStruct(typ reflect.Type) *ReflectedStruct {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}
	if s, ok := structTypes[typ]; ok {
		return s
	}
	s := &ReflectedStruct{
		Typ:        typ,
		Methods:    map[string]int{},
		Fields:     map[string][]int{},
		Construtor: -1,
	}
	structTypes[typ] = s
	var each func(typ reflect.Type, ix []int)
	each = func(typ reflect.Type, ix []int) {
		for typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		for l, i := typ.NumField(), 0; i < l; i++ {
			f := typ.Field(i)
			if f.Anonymous {
				each(f.Type, append(ix, f.Index...))
			}
			s.Fields[f.Name] = append(ix, f.Index...)
		}
	}
	each(typ, nil)
	typPtr := reflect.PtrTo(typ)
	for l, i := typPtr.NumMethod(), 0; i < l; i++ {
		m := typPtr.Method(i)
		s.Methods[m.Name] = i
	}

	if ix, ok := s.Methods["Constructor"]; ok {
		delete(s.Methods, "Constructor")
		s.Construtor = ix
	}

	return s
}

func (s *ReflectedStruct) Fqn() string {
	return s.Typ.PkgPath() + "." + s.Typ.Name()
}

func (ReflectedStruct) Name() string {
	return "reflect-struct"
}

func (s *ReflectedStruct) String() string {
	return fmt.Sprintf("<reflect-struct %s>", s.Fqn())
}

func (s *ReflectedStruct) Call(args ...Object) (ret Object, err error) {
	obj := &ReflectedStructInstance{Struct: s, Instance: reflect.New(s.Typ)}
	if s.Construtor != -1 {
		m := &ReflectedStructInstanceMethod{
			MethodName: "Constructor",
			Method:     obj.Instance.Method(s.Construtor),
			Instance:   obj,
		}
		if _, err = m.Call(args...); err != nil {
			return
		}
	}
	return obj, nil
}

func (s *ReflectedStruct) CanCall() bool {
	return true
}

func (s *ReflectedStruct) CanCallContext() bool {
	return true
}

type ReflectedStructInstance struct {
	ObjectImpl
	Struct   *ReflectedStruct
	Instance reflect.Value
}

func (ReflectedStructInstance) Name() string {
	return "reflect-struct-instance"
}

func (s *ReflectedStructInstance) String() string {
	return fmt.Sprintf("<reflect-struct-instance %s: %s>", s.Struct.Fqn(), fmt.Sprint(s.Instance.Interface()))
}

func (i *ReflectedStructInstance) Interface() interface{} {
	return i.Instance.Addr().Interface()
}

func (i *ReflectedStructInstance) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := ToString(index)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	var ix int
	if ix, ok = i.Struct.Methods[strIdx]; !ok {
		ix, ok := i.Struct.Fields[strIdx]
		if !ok {
			res = UndefinedValue
			return
		}
		value := i.Instance.Elem().FieldByIndex(ix)
		if value.Kind() == reflect.Struct {
			value = value.Addr()
		}
		return FromInterface(value)
	} else {
		m := i.Instance.Method(ix)
		res = &ReflectedStructInstanceMethod{
			Method:     m,
			Instance:   i,
			MethodName: strIdx,
		}
	}
	return
}

func (i *ReflectedStructInstance) IndexSet(index, value Object) error {
	strIdx, ok := ToString(index)
	if !ok {
		return ErrInvalidIndexType
	}
	ix, ok := i.Struct.Fields[strIdx]
	if !ok {
		return fmt.Errorf("field %q does not exists in %s", strIdx, i.Struct.Fqn())
	}
	f := i.Instance.Elem().FieldByIndex(ix)
	v := reflect.ValueOf(ToInterface(value))
	if v.Type().ConvertibleTo(f.Type()) {
		f.Set(v.Convert(f.Type()))
	} else if v.Type().AssignableTo(f.Type()) {
		f.Set(v)
	} else {
		return ErrInvalidArgumentType{
			Name:     "Value",
			Expected: "go:" + f.Type().String() + "(compatible)",
			Found:    value.TypeName() + " as go:" + v.Type().String(),
		}
	}
	return nil
}

type ReflectedStructInstanceMethod struct {
	ObjectImpl
	Instance   *ReflectedStructInstance
	MethodName string
	Method     reflect.Value
}

func (ReflectedStructInstanceMethod) Name() string {
	return "reflect-struct-method"
}

func (s *ReflectedStructInstanceMethod) String() string {
	return fmt.Sprintf("<reflect-struct-method %s#%s>", s.Instance.Struct.Fqn(), s.MethodName)
}

func (s *ReflectedStructInstanceMethod) CanCall() bool {
	return true
}

func (s *ReflectedStructInstanceMethod) CanCallContext() bool {
	return true
}

func (s *ReflectedStructInstanceMethod) Call(args ...Object) (ret Object, err error) {
	var (
		ctx   *Context
		i     = 0
		rargs []reflect.Value
	)

	if s.Method.Type().NumIn() > 0 && s.Method.Type().In(0).AssignableTo(reflect.TypeOf((*Context)(nil))) {
		ctx = args[0].(*Context)
		args = args[1:]
		rargs = make([]reflect.Value, len(args)+1)
		rargs[0] = reflect.ValueOf(ctx)
		i++
	} else {
		args = args[1:]
		rargs = make([]reflect.Value, len(args))
	}

	for _, v := range args {
		vi := ToInterface(v)
		if vi != nil {
			rargs[i] = reflect.ValueOf(vi)
		}
		i++
	}
	out := s.Method.Call(rargs)
	l := len(out)
	if l > 0 {
		errt := reflect.TypeOf((*error)(nil)).Elem()
		if out[l-1].Type().AssignableTo(errt) {
			if err = out[l-1].Interface().(error); err != nil {
				return
			}
			out = out[1:]
			l--
		}
	}
	if l == 0 {
		ret = UndefinedValue
		return
	} else if l == 1 {
		return FromInterface(out[0])
	}
	retArray := &Array{Value: make([]Object, l)}
	for i, out := range out {
		if ret, err = FromInterface(out); err != nil {
			return
		}
		retArray.Value[i] = ret
	}
	return retArray, nil
}

func FromReflectValue(value reflect.Value) (res Object, err error) {
	s := NewReflectedStruct(reflect.Indirect(value).Type())
	if s == nil {
		return
	}
	err = nil
	res = &ReflectedStructInstance{Struct: s, Instance: value}
	return
}

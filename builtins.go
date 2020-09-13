package tengo

import (
	"context"
	"time"
)

var builtinFuncs = []*BuiltinContextFunction{
	{
		Name:  "len",
		Value: builtinLen,
	},
	{
		Name:  "copy",
		Value: builtinCopy,
	},
	{
		Name:  "append",
		Value: builtinAppend,
	},
	{
		Name:  "delete",
		Value: builtinDelete,
	},
	{
		Name:  "splice",
		Value: builtinSplice,
	},
	{
		Name:  "string",
		Value: builtinString,
	},
	{
		Name:  "int",
		Value: builtinInt,
	},
	{
		Name:  "bool",
		Value: builtinBool,
	},
	{
		Name:  "float",
		Value: builtinFloat,
	},
	{
		Name:  "char",
		Value: builtinChar,
	},
	{
		Name:  "bytes",
		Value: builtinBytes,
	},
	{
		Name:  "time",
		Value: builtinTime,
	},
	{
		Name:  "is_int",
		Value: builtinIsInt,
	},
	{
		Name:  "is_float",
		Value: builtinIsFloat,
	},
	{
		Name:  "is_string",
		Value: builtinIsString,
	},
	{
		Name:  "is_bool",
		Value: builtinIsBool,
	},
	{
		Name:  "is_char",
		Value: builtinIsChar,
	},
	{
		Name:  "is_bytes",
		Value: builtinIsBytes,
	},
	{
		Name:  "is_array",
		Value: builtinIsArray,
	},
	{
		Name:  "is_immutable_array",
		Value: builtinIsImmutableArray,
	},
	{
		Name:  "is_map",
		Value: builtinIsMap,
	},
	{
		Name:  "is_immutable_map",
		Value: builtinIsImmutableMap,
	},
	{
		Name:  "is_iterable",
		Value: builtinIsIterable,
	},
	{
		Name:  "is_time",
		Value: builtinIsTime,
	},
	{
		Name:  "is_error",
		Value: builtinIsError,
	},
	{
		Name:  "is_undefined",
		Value: builtinIsUndefined,
	},
	{
		Name:  "is_function",
		Value: builtinIsFunction,
	},
	{
		Name:  "is_callable",
		Value: builtinIsCallable,
	},
	{
		Name:  "type_name",
		Value: builtinTypeName,
	},
	{
		Name:  "format",
		Value: builtinFormat,
	},
	{
		Name:  "context",
		Value: builtinContext,
	},
	{
		Name:  "context_timeout",
		Value: builtinContextTimeout,
	},
	{
		Name:  "context_deadline",
		Value: builtinContextWithDeadline,
	},
	{
		Name:  "context_canceler",
		Value: builtinContextCanceler,
	},
	{
		Name:  "context_cancel",
		Value: builtinContextCancel,
	},
	{
		Name:  "struct",
		Value: builtinStruct,
	},
	{
		Name:  "new",
		Value: builtinNew,
	},
}

// GetAllBuiltinFunctions returns all builtin function objects.
func GetAllBuiltinFunctions() []*BuiltinContextFunction {
	return append([]*BuiltinContextFunction{}, builtinFuncs...)
}

func builtinTypeName(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return &String{Value: args[0].TypeName()}, nil
}

func builtinIsString(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*String); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsInt(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFloat(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBool(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsChar(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBytes(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bytes); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsArray(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Array); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsImmutableArray(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*ImmutableArray); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsMap(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Map); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsImmutableMap(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*ImmutableMap); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsTime(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Time); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsError(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Error); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsUndefined(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0] == UndefinedValue {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFunction(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch args[0].(type) {
	case *CompiledFunction:
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsCallable(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanCall() {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsIterable(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanIterate() {
		return TrueValue, nil
	}
	return FalseValue, nil
}

// len(obj object) => int
func builtinLen(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *ImmutableArray:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *String:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Bytes:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Map:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *ImmutableMap:
		return &Int{Value: int64(len(arg.Value))}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array/string/bytes/map",
			Found:    arg.TypeName(),
		}
	}
}

func builtinFormat(_ *Context, args ...Object) (Object, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, ErrWrongNumArguments
	}
	format, ok := args[0].(*String)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &String{Value: s}, nil
}

func builtinCopy(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return args[0].Copy(), nil
}

func builtinString(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*String); ok {
		return args[0], nil
	}
	v, ok := ToString(args[0])
	if ok {
		if len(v) > MaxStringLen {
			return nil, ErrStringLimit
		}
		return &String{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinInt(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return args[0], nil
	}
	v, ok := ToInt64(args[0])
	if ok {
		return &Int{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinFloat(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return args[0], nil
	}
	v, ok := ToFloat64(args[0])
	if ok {
		return &Float{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinBool(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return args[0], nil
	}
	v, ok := ToBool(args[0])
	if ok {
		if v {
			return TrueValue, nil
		}
		return FalseValue, nil
	}
	return UndefinedValue, nil
}

func builtinChar(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return args[0], nil
	}
	v, ok := ToRune(args[0])
	if ok {
		return &Char{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinBytes(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}

	// bytes(N) => create a new bytes with given size N
	if n, ok := args[0].(*Int); ok {
		if n.Value > int64(MaxBytesLen) {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: make([]byte, int(n.Value))}, nil
	}
	v, ok := ToByteSlice(args[0])
	if ok {
		if len(v) > MaxBytesLen {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

func builtinTime(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Time); ok {
		return args[0], nil
	}
	v, ok := ToTime(args[0])
	if ok {
		return &Time{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return UndefinedValue, nil
}

// append(arr, items...)
func builtinAppend(_ *Context, args ...Object) (Object, error) {
	if len(args) < 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Array{Value: append(arg.Value, args[1:]...)}, nil
	case *ImmutableArray:
		return &Array{Value: append(arg.Value, args[1:]...)}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    arg.TypeName(),
		}
	}
}

// builtinDelete deletes Map keys
// usage: delete(map, "key")
// key must be a string
func builtinDelete(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Map:
		if key, ok := args[1].(*String); ok {
			delete(arg.Value, key.Value)
			return UndefinedValue, nil
		}
		return nil, ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map",
			Found:    arg.TypeName(),
		}
	}
}

// builtinSplice deletes and changes given Array, returns deleted items.
// usage:
// deleted_items := splice(array[,start[,delete_count[,item1[,item2[,...]]]])
func builtinSplice(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return nil, ErrWrongNumArguments
	}

	array, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}
	arrayLen := len(array.Value)

	var startIdx int
	if argsLen > 1 {
		arg1, ok := args[1].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int",
				Found:    args[1].TypeName(),
			}
		}
		startIdx = int(arg1.Value)
		if startIdx < 0 || startIdx > arrayLen {
			return nil, ErrIndexOutOfBounds
		}
	}

	delCount := len(array.Value)
	if argsLen > 2 {
		arg2, ok := args[2].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int",
				Found:    args[2].TypeName(),
			}
		}
		delCount = int(arg2.Value)
		if delCount < 0 {
			return nil, ErrIndexOutOfBounds
		}
	}
	// if count of to be deleted items is bigger than expected, truncate it
	if startIdx+delCount > arrayLen {
		delCount = arrayLen - startIdx
	}
	// delete items
	endIdx := startIdx + delCount
	deleted := append([]Object{}, array.Value[startIdx:endIdx]...)

	head := array.Value[:startIdx]
	var items []Object
	if argsLen > 3 {
		items = make([]Object, 0, argsLen-3)
		for i := 3; i < argsLen; i++ {
			items = append(items, args[i])
		}
	}
	items = append(items, array.Value[endIdx:]...)
	array.Value = append(head, items...)

	// return deleted items
	return &Array{Value: deleted}, nil
}

// builtinContext returns context or VM context copy with values.
// usage:
// vm_context := context()
// background_context := context(undefined, [key1,value1[,keyN,valueN[,...]]])
// context_with_values := context(context_var, key1,value1[,keyN,valueN[,...]])
func builtinContext(vmCtx *Context, args ...Object) (Object, error) {
	var (
		ctx *Context
		ok  bool
		l   = len(args)
	)
	if l == 0 {
		return vmCtx, nil
	} else if ctx, ok = args[0].(*Context); ok {
		args = args[1:]
		l--
	} else if args[0] == UndefinedValue {
		ctx = &Context{Value: context.Background()}
		args = args[1:]
		l--
	} else if l%2 != 0 {
		return nil, ErrWrongNumArguments
	}
	if l == 0 {
		return ctx, nil
	}

	ctx = &Context{Value: ctx.Value}
	for i := 0; i < l; i += 2 {
		ctx.Value = context.WithValue(ctx.Value, ToInterface(args[i]), args[i+1])
	}
	return ctx, nil
}

// builtinContextWithDeadline returns VM context copy with timeout.
// if context is undefined, returns timeout of `context.Background()`.
// usage:
// new_context := context_deadline(context, time)
func builtinContextWithDeadline(_ *Context, args ...Object) (Object, error) {
	var ctx *Context
	if l := len(args); l != 1 {
		return nil, ErrWrongNumArguments
	} else if args[0] == UndefinedValue {
		ctx = &Context{Value: context.Background()}
	} else {
		var ok bool
		if ctx, ok = args[0].(*Context); !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "context",
				Found:    args[0].TypeName(),
			}
		}
		ctx = ctx.Copy().(*Context)
	}

	t1, ok := ToTime(args[1])
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	ctx = ctx.Copy().(*Context)
	ctx.Value, ctx.cancel = context.WithDeadline(ctx.Value, t1)
	return ctx, nil
}

// builtinContextTimeout returns VM context copy with timeout.
// if context is undefined, returns timeout of `context.Background()`.
// usage:
// new_context := context_timeout(context, duration)
func builtinContextTimeout(_ *Context, args ...Object) (Object, error) {
	var ctx *Context
	if l := len(args); l != 1 {
		return nil, ErrWrongNumArguments
	} else if args[0] == UndefinedValue {
		ctx = &Context{Value: context.Background()}
	} else {
		var ok bool
		if ctx, ok = args[0].(*Context); !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "context",
				Found:    args[0].TypeName(),
			}
		}
		ctx = ctx.Copy().(*Context)
	}

	var dur time.Duration

	switch v := args[1].(type) {
	case *Int:
		dur = time.Duration(v.Value)
	case *Float:
		dur = time.Duration(v.Value)
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int|float",
			Found:    args[0].TypeName(),
		}
	}

	dur *= time.Second

	ctx = ctx.Copy().(*Context)
	ctx.Value, ctx.cancel = context.WithTimeout(ctx.Value, dur)
	return ctx, nil
}

// builtinContextCanceler returns VM context copy with canceler.
// if not have args, returns canceler of vm context.
// if context is undefined, returns canceler of `context.Background()`.
// usage:
// new_context := context_canceler([context])
// context_cancel(new_context)
func builtinContextCanceler(vmCtx *Context, args ...Object) (Object, error) {
	var ctx *Context
	if l := len(args); l == 0 {
		ctx = &Context{Value: vmCtx.Value}
	} else if l != 1 {
		return nil, ErrWrongNumArguments
	} else if args[0] == UndefinedValue {
		ctx = &Context{Value: context.Background()}
	} else {
		var ok bool
		if ctx, ok = args[0].(*Context); !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "context",
				Found:    args[0].TypeName(),
			}
		}
		ctx = ctx.Copy().(*Context)
	}
	ctx.Value, ctx.cancel = context.WithCancel(ctx.Value)
	return ctx, nil
}

// builtinContextCancel cancel context.
// usage:
// new_context := context_canceler(context)
// context_cancel(new_context)
func builtinContextCancel(_ *Context, args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	var ctx *Context
	if args[0] == UndefinedValue {
		ctx = &Context{Value: context.Background()}
	} else {
		var ok bool
		if ctx, ok = args[0].(*Context); !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "context",
				Found:    args[0].TypeName(),
			}
		}
	}

	if ctx.cancel == nil {
		return nil, ErrContextNotCancelable
	}
	ctx.cancel()
	// set fake canceler
	ctx.cancel = func() {}
	return UndefinedValue, nil
}

// builtinStruct define new struct from map
// usage:
// struct({
//   fields: {
//		count: 0
//   },
//	 init: func(this) {
//		this.count = 1
//   }
//	 funcs: {
//		increment: func(this, value) {
//			return this.count++
//		}
//	 }
// })
func builtinStruct(_ *Context, args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen < 1 {
		return nil, ErrWrongNumArguments
	}

	if m, ok := args[0].(*Map); !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map",
			Found:    args[0].TypeName(),
		}
	} else {
		var (
			fieldsM, funcsM *Map
		)
		if fields, ok := m.Value["fields"]; ok {
			if fieldsM, ok = fields.(*Map); !ok {
				return nil, ErrInvalidMapIndexValueType{
					MapName:   "first arg",
					IndexName: "fields",
					Expected:  "map",
					Found:     fields.TypeName(),
				}
			}
		}
		if funcs, ok := m.Value["funcs"]; ok {
			if funcsM, ok = funcs.(*Map); !ok {
				return nil, ErrInvalidMapIndexValueType{
					MapName:   "first arg",
					IndexName: "funcs",
					Expected:  "map",
					Found:     funcs.TypeName(),
				}
			}
		}
		return NewStruct(fieldsM, funcsM)
	}
}

// builtinStruct define new struct from map
// usage:
// struct({
//   fields: {
//		count: 0
//   },
//	 init: func(this) {
//		this.count = 1
//   }
//	 funcs: {
//		increment: func(this, value) {
//			return this.count++
//		}
//	 }
// })
func builtinNew(ctx *Context, args ...Object) (res Object, err error) {
	argsLen := len(args)
	if argsLen < 1 {
		return nil, ErrWrongNumArguments
	}
	var (
		typ    = args[0]
		fields *Map
	)
	args = args[1:]

	if len(args) > 0 {
		var ok bool
		if fields, ok = args[0].(*Map); !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "second",
				Expected: "map",
				Found:    args[0].TypeName(),
			}
		}
		args = args[1:]
	}
	switch t := typ.(type) {
	case *Struct, *ReflectedStruct:
		if res, err = t.Call(append([]Object{ctx}, args...)...); err != nil {
			return nil, err
		}
		if fields != nil {
			for key, value := range fields.Value {
				if err = res.IndexSet(&String{Value: key}, value); err != nil {
					return nil, err
				}
			}
		}
		return
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "struct|reflect-struct",
			Found:    t.TypeName(),
		}
	}
}

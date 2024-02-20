package evaluator

import (
	"macaque/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of argument, got=%d, expected 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("wrong argument type for 'len()': %s", arg.Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of argument, got=%d, expected 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("wrong argument type for 'first()': %s, expected 'ARRAY'", args[0].Type())
			}

			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of argument, got=%d, expected 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("wrong argument type for 'last()': %s, expected 'ARRAY'", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				return array.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of argument, got=%d, expected 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("wrong argument type for 'rest()': %s, expected 'ARRAY'", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				newElems := make([]object.Object, length-1, length-1)
				copy(newElems, array.Elements[1:length])

				return &object.Array{Elements: newElems}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of argument, got=%d, expected 2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("wrong argument type for 'push()': %s, expected 'ARRAY'", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)

			newElems := make([]object.Object, length+1, length+1)
			copy(newElems, array.Elements)
			newElems[length] = args[1]

			return &object.Array{Elements: newElems}
		},
	},
}

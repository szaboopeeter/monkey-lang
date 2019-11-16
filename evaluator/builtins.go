package evaluator

import "monkey-lang/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Invalid amount of arguments. Expected=%d, got=%d", 1, len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("Invalid argument passed to `len()`. Got=%s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Invalid amount of arguments. Expected=%d, got=%d", 1, len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Invalid argument passed to `first()`. Expected=ARRAY, got=%s", args[0].Type())
			}

			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Invalid amount of arguments. Expected=%d, got=%d", 1, len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Invalid argument passed to `last()`. Expected=ARRAY, got=%s", args[0].Type())
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
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Invalid amount of arguments. Expected=%d, got=%d", 1, len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Invalid argument passed to `rest()`. Expected=ARRAY, got=%s", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, array.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Function: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Invalid amount of arguments. Expected=%d, got=%d", 2, len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Invalid argument passed to `push()`. Expected=ARRAY, got=%s", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, array.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
}

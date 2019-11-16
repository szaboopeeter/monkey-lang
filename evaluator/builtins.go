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
			default:
				return newError("Invalid argument passed to `len()`. Got=%s", args[0].Type())
			}
		},
	},
}

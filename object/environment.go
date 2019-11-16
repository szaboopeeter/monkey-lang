package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (environment *Environment) Get(name string) (Object, bool) {
	object, ok := environment.store[name]

	if !ok && environment.outer != nil {
		object, ok = environment.outer.Get(name)
	}

	return object, ok
}

func (enviornment *Environment) Set(name string, value Object) Object {
	enviornment.store[name] = value
	return value
}

package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (environment *Environment) Get(name string) (Object, bool) {
	object, ok := environment.store[name]
	return object, ok
}

func (enviornment *Environment) Set(name string, value Object) Object {
	enviornment.store[name] = value
	return value
}

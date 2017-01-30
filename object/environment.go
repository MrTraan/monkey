package object

import "fmt"

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Upsert(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Environment) Set(name string, val Object) Object {
	if _, ok := e.store[name]; ok {
		return &Error{Message: fmt.Sprintf("Identifier '%s' has already been declared", name)}
	}
	e.store[name] = val
	return val
}

func (e *Environment) Update(name string, val Object) Object {
	if _, ok := e.store[name]; !ok {
		return &Error{Message: fmt.Sprintf("%s is not defined", name)}
	}
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

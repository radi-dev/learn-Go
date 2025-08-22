package pointrserrs

import "fmt"

type Entity struct {
	Name       string
	Function   string
	GreetCount int
}

func (ent Entity) Greet() (out string) {
	out = fmt.Sprintf("Hi, I am %s, an Entity that %s", ent.Name, ent.Function)
	ent.GreetCount++ //Does nothing. ent is not a pointer, no mutation.
	return
}

func (ent *Entity) GreetMut() (out string) {
	out = fmt.Sprintf("Hi, I am %s, an Entity that %s", ent.Name, ent.Function)
	ent.GreetCount++
	return
}

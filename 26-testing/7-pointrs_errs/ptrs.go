package pointrserrs

import "fmt"

type Count int

type Entity struct {
	Name       string
	Function   string
	GreetCount Count
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

func (ent *Entity) Vex() (out string) {
	out = fmt.Sprintf("Hi, I am %s, an Entity that %s", ent.Name, ent.Function)
	ent.GreetCount--
	return
}

// ============== Extras

func (c Count) Stringify() string {
	return fmt.Sprint(c)
}

package structsmethdsintrfcs

import "fmt"

type Entity struct {
	Name     string
	Function string
}

type Agent interface {
	Greet() string
	// Chat()  // Uncommenting this would cause compilation error in the test because Entity doesn't have Chat method defined
}

func Greet1(ent Entity) (out string) {
	out = fmt.Sprintf("Hi, I am %s, an Entity that %s", ent.Name, ent.Function)
	return
}

func (ent Entity) Greet() (out string) {
	out = fmt.Sprintf("Hi, I am %s, an Entity that %s", ent.Name, ent.Function)
	return
}

// ========= Anonymous structs =====

var Prof = struct {
	Name     string
	Function string
}{"Professor", "Teaches students"}

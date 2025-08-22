package structsmethdsintrfcs

import (
	"fmt"
	"testing"
	"testing_test/utils"
)

func TestStructs(t *testing.T) {
	entity := Entity{"Angel", "Gives daily motivation"}
	got := Greet1(entity)
	want := fmt.Sprintf("Hi, I am %s, an Entity that %s", entity.Name, entity.Function)

	utils.AssertCorrectMessage(t, got, want)
}

func TestMethod(t *testing.T) {
	entity := Entity{"Angel", "Gives daily motivation"}
	got := entity.Greet()
	want := fmt.Sprintf("Hi, I am %s, an Entity that %s", entity.Name, entity.Function)

	utils.AssertCorrectMessage(t, got, want)
}

func TestInterface(t *testing.T) {
	BasicGreet := func(a Agent) string {
		return fmt.Sprintf("This is a greeting message: %s", a.Greet())
	}
	entity := Entity{"Angel", "Gives daily motivation"}
	got := BasicGreet(entity)
	want := fmt.Sprintf("This is a greeting message: Hi, I am %s, an Entity that %s", entity.Name, entity.Function)
	t.Run("Test that entity is Agent", func(t *testing.T) { utils.AssertCorrectMessage(t, got, want) })

}

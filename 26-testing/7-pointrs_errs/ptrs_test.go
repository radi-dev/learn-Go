package pointrserrs

import (
	"testing"
	"testing_test/utils"
)

func TestGreetMetod(t *testing.T) {
	ent1 := Entity{Name: "Symposium", Function: "Creates musical notes"}
	t.Run("Test that Greet method does not modify GreetCount", func(t *testing.T) {
		for range 10 {
			ent1.Greet()
		}
		got := ent1.GreetCount
		var want Count = 0

		utils.AssertCorrectMessage(t, got, want)
	})

	t.Run("Test that GreGreetMutet method modifies GreetCount", func(t *testing.T) {
		for range 10 {
			ent1.GreetMut()
		}
		got := ent1.GreetCount
		want := Count(10)

		utils.AssertCorrectMessage(t, got, want)
	})

}

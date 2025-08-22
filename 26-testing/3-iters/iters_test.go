package iters

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	fmt.Println(Repeat("he", 3))
	// Output: hehehe
}

func TestRepeat(t *testing.T) {
	t.Run("repeat", func(t *testing.T) {
		got := Repeat("he", 3)
		want := "hehehe"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

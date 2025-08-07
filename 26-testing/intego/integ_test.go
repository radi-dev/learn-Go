package intego

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

func TestAdd(t *testing.T) {
	t.Run("Test Addition", func(t *testing.T) {
		got := Add(2, 3)
		want := 5
		if got != want {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})
}

package benchmrk

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

func BenchmarkEfficientRepeat(b *testing.B) {
	for b.Loop() {
		EfficientRepeat("a", 10)
	}
}

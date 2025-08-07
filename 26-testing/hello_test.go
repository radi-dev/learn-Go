package testing_test

import (
	"reflect"
	"testing"
)

func HelloTest(t *testing.T) {
	got := Hello()
	want := "Hello, testing ground"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithArg(t *testing.T) {
	got := Hello2("Ra")
	want := "Hello, testing ground...Ra"

	gotKind := reflect.ValueOf(want).Kind()
	reqKind := reflect.String

	if gotKind != reqKind {
		t.Errorf("Req kind: %s, got: %s", reqKind, gotKind)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

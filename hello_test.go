package helloversioning

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world v1."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

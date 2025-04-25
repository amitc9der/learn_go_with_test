package depedencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Amit")

	got := buffer.String()
	want := "Hello, Amit"

	if got != want {
		t.Errorf("got: %s, wanted: %q", got, want)
	}
}

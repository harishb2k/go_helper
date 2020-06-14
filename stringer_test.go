package helper

import "testing"

func TestHello(t *testing.T) {
	output, _ := Stringify("Test")
	if output != "Test*-" {
		t.Errorf("Stringify() = %q, want %q", output, "Test*")
	}
}

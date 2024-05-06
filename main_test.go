package main

import "testing"

func TestValidateJson(t *testing.T) {
	got := validateJson("tests/step1/valid.json")
	want := 0

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

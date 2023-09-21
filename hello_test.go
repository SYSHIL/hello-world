package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("ilhan")
	want := "ILHAN"
	if got != want {
		t.Errorf("got:%q want:%q", got, want)
	}
}

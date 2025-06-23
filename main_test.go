package main

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

func TestHello(t *testing.T) {
	want := "Hello Go"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

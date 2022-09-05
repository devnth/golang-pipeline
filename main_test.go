package main

import "testing"

func TestHello(t *testing.T) {

	want := "hello golang"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

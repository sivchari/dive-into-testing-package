package main

import "testing"

func TestAdd(t *testing.T) {
	i, j, want := 1, 1, 2
	got := Add(i, j)
	if want != got {
	}
}

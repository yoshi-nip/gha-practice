package main

import "testing"

func TestEvenOOdd(t *testing.T) {
	result := EvenOOdd(10)
	if result != "even"{
		t.Errorf("expected: even, acutual: %s",result)
	}
}

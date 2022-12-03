package main

import (
	"reflect"
	"testing"
)

func Test_winningShape(t *testing.T) {
	assertEqual(t, pickWinner(A, B), B)
	assertEqual(t, pickWinner(B, C), C)
	assertEqual(t, pickWinner(C, A), A)
}

func Test_losingShape(t *testing.T) {
	assertEqual(t, losingShape(B), A)
	assertEqual(t, losingShape(C), B)
	assertEqual(t, losingShape(A), C)
}

func assertEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Expected to be equal: %v, %v", a, b)
	}
}

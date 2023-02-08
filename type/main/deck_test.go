package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 51 {
		t.Errorf("Expected: 51, Actual:%v", len(d))
	}
}

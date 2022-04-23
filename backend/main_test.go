package main

import "testing"

func TestAdd(t *testing.T) {
	if 1 == 2 {
		t.Errorf("Add test failed")
	}
}
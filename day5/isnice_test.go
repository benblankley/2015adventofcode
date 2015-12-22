package main

import (
	"testing"
	)

func TestIsnice(t *testing.T) {
	expected := true
	actual := isnice("antonio")


	if actual != expected {
		t.Errorf("isnice (antonio) actual: %v, expected %v", actual, expected)
	}
}

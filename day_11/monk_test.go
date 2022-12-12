package main

import (
	"testing"
)

func TestCountVisitsTail9(t *testing.T){
	input := buildMonke()
	want:= 98280
	actual:=calculateMonkeyBusiness(input)
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted %v`, input, actual, want)
	}
}
package main

import (
	"testing"
	// "fmt"
)
func TestSumWithZero(t *testing.T){
	input := "1=-0-2"
	actual := sum(parseInput(input))
	want := "1=-0-2"
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted sum to be %v`, input, actual, want )
	}
}
func TestSumTwoNumbers(t *testing.T){
	input := "1=-0-2\n12111"
	actual := sum(parseInput(input))
	want := "1-111="//2653
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted sum to be %v`, input, actual, want )
	}
}
func TestSumThreeNumbers(t *testing.T){
	input := "1=-0-2\n12111\n2=0="
	actual := sum(parseInput(input))
	want := "10=-01"//2851
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted sum to be %v`, input, actual, want )
	}
}

func TestSumAll(t *testing.T){
	input := "1=-0-2\n12111\n2=0=\n21\n2=01\n111\n20012\n112\n1=-1=\n1-12\n12\n1=\n122"
	actual := sum(parseInput(input))
	want := "2=-1=0"
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted sum to be %v`, input, actual, want )
	}
}
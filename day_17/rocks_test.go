package main

import (
	"testing"
)

func TestFirstFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 1)
	want := 1
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestSecondFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 2)
	want := 4
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestThirdFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 3)
	want := 6
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestFourthFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 4)
	want := 7
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestFifthFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 5)
	want := 9
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestSixthFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 6)
	want := 10
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestSEventhFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 7)
	want := 13
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestEighthFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 8)
	want := 15
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestNinethFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 200, 7, 9)
	want := 17
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
func TestAllFall(t *testing.T){
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	actual:= rocksFall(input, 4000, 7, 2022)
	want := 3068
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted height to be %v`, input, actual, want )
	}
}
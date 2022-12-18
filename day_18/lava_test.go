package main

import (
	"testing"
)

func TestOneBlock(t *testing.T){
	input := "1,1,1"
	positions:= getPositions(input)
	blob := makeBlob(positions)
	actual:= countSides(blob, positions)
	want := 6
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted sides to be %v`, input, actual, want )
	}
}
func TestOneBlock0(t *testing.T){
	input := "0,1,1"
	positions:= getPositions(input)
	blob := makeBlob(positions)
	actual:= countSides(blob, positions)
	want := 6
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted sides to be %v`, input, actual, want )
	}
}
func TestTwoBlocks(t *testing.T){
	input := "1,1,1\n2,1,1"
	positions:= getPositions(input)
	blob := makeBlob(positions)
	actual:= countSides(blob, positions)
	want := 10
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted sides to be %v`, input, actual, want )
	}
}
func TestAllBlocks(t *testing.T){
	input := "2,2,2\n1,2,2\n3,2,2\n2,1,2\n2,3,2\n2,2,1\n2,2,3\n2,2,4\n2,2,6\n1,2,5\n3,2,5\n2,1,5\n2,3,5"
	positions:= getPositions(input)
	blob := makeBlob(positions)
	actual:= countSides(blob, positions)
	want := 64
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted sides to be %v`, input, actual, want )
	}
}

func TestAllBlocksNoInnder(t *testing.T){
	input := "2,2,2\n1,2,2\n3,2,2\n2,1,2\n2,3,2\n2,2,1\n2,2,3\n2,2,4\n2,2,6\n1,2,5\n3,2,5\n2,1,5\n2,3,5"
	positions := getPositions(input)
	blob := makeBlob(positions)
	actual:= countSidesNoInner(blob, positions)
	want := 58
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted sides to be %v`, input, actual, want )
	}
}
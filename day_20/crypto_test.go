package main

import (
	"testing"
)

func TestSum(t *testing.T){
	input:= "1\n2\n-3\n3\n-2\n0\n4"
	list, zero := createNodes(parseLines(input))
	setNeighbors(list)
	unmixAll(list)
	actual:= sum(zero)
	want := 3

		if want != actual {
			t.Fatalf(`Given %v, got %v, but wanted unmix to be %v`, input, actual, want )
		}

}

func TestSum2(t *testing.T){
	input:= "1\n2\n-3\n3\n-2\n0\n4"
	list, zero := createNodes(parseLines(input))
	decrypt(list, 811589153)
	setNeighbors(list)
	superUnmixAll(list)
	actual:= sum(zero)
	want := 1623178306

		if want != actual {
			t.Fatalf(`Given %v, got %v, but wanted unmix to be %v`, input, actual, want )
		}

}
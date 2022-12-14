package main

import (
	"testing"
)

func TestBuildCave(t *testing.T){
	input := "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"
	want:= 20
	cave := buildCave(parseLines(input))
	var actual int
	for y, row := range cave {
		for x, _ := range row {

			if cave[y][x] == '#' {
				actual+=1
			}
		}
	}
	t.Log(cave)
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted length to be %v`, input, actual, want )
	}
}
func TestLandslide(t *testing.T){
	input := "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"
	want:= 93
	cave := buildCave(parseLines(input))
	filledCave, actual := landslide(cave, Position{500,0})
	t.Log(filledCave)
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted length to be %v`, input, actual, want )
	}
}
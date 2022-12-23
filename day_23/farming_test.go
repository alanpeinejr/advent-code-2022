package main

import (
	"testing"
	// "fmt"
)

func TestConsiderMoves(t *testing.T){
	input := ".....\n..##.\n..#..\n.....\n..##.\n....."
	plot, elves := parseMap(input, 5)
	for i, elf := range elves {
		elves[i] = considerMoves(elf, plot, []Heading{north, south, west, east})

	}
	wants := []Heading{north, north, south, north, north}
	for i, want := range wants {
		if want != elves[i].heading{
			t.Fatalf(`Given %v, got %v, but wanted items to be %v`, input, elves[i].heading, want )
		}
	}
}
func TestConsiderMoves2(t *testing.T){
	input := "..##.\n.....\n..#..\n...#.\n..#..\n....."
	plot, elves := parseMap(input, 5)
	for i, elf := range elves {
		elves[i] = considerMoves(elf, plot, []Heading{south, west, east, north})

	}
	wants := []Heading{south, south, west, east, south}
	for i, want := range wants {
		if want != elves[i].heading{
			t.Fatalf(`Given %v, got %v, but wanted items to be %v`, input, elves[i].heading, want )
		}
	}
}
func TestConsiderMoves3(t *testing.T){
	input := ".....\n..##.\n.#...\n....#\n.....\n..#.."
	plot, elves := parseMap(input, 5)
	for i, elf := range elves {
		elves[i] = considerMoves(elf, plot, []Heading{west, east, north, south})

	}
	wants := []Heading{north, east, west, stationary, stationary}
	for i, want := range wants {
		if want != elves[i].heading{
			t.Fatalf(`Given %v, got %v, but wanted items to be %v`, input, elves[i].heading, want )
		}
	}
}
func TestConsiderMoves4(t *testing.T){
	input := "..#..\n....#\n#....\n....#\n.....\n..#.."
	plot, elves := parseMap(input, 5)
	for i, elf := range elves {
		elves[i] = considerMoves(elf, plot, []Heading{east, north, south, west})

	}
	wants := []Heading{stationary, stationary, stationary, stationary, stationary}
	for i, want := range wants {
		if want != elves[i].heading{
			t.Fatalf(`Given %v, got %v, but wanted items to be %v`, input, elves[i].heading, want )
		}
	}
}
func TestRound1(t *testing.T){
	input := ".....\n..##.\n..#..\n.....\n..##.\n....."
	plot, elves := parseMap(input, 5)
	plot, elves, _ = round(plot, elves, []Heading{north, south, west, east})
	minX, maxX, minY, maxY := findRectangle(elves)
	actual := countGround(plot, minX, maxX, minY, maxY)
	want :=5
		if want != actual{
			t.Fatalf(`Given %v, got %v, but wanted items to be %v`, input, actual, want )
		}
}
func TestRound2(t *testing.T){
	input := ".....\n..##.\n..#..\n.....\n..##.\n....."
	plot, elves := parseMap(input, 5)
	plot, elves, _ = round(plot, elves, []Heading{north, south, west, east})
	plot, elves, _ = round(plot, elves, []Heading{south, west, east, north})
	minX, maxX, minY, maxY := findRectangle(elves)
	actual := countGround(plot, minX, maxX, minY, maxY)
	want :=15
		if want != actual{
			t.Fatalf(`Given %v, got %v, but wanted items to be %v`, input, actual, want )
		}
}
func TestUntilNoMoves(t *testing.T){
	input := ".....\n..##.\n..#..\n.....\n..##.\n....."
	plot, elves := parseMap(input, 5)
	actual := untilNoMoves(plot, elves)
	want :=4
		if want != actual{
			t.Fatalf(`Given %v, got %v, but wanted items to be %v`, input, actual, want )
		}
}
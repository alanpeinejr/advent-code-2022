package main

import (
	"testing"
)

func TestLastCapitalism(t *testing.T){
	input := "Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian."
	meansOfProduction := parseLines(input)
	qualityScore := capitalism(meansOfProduction[0], 24)
	want:= 9
	if want != qualityScore {
		t.Fatalf(`Given %v, got %v, but score to be %v`, input, qualityScore, want )
	}

}
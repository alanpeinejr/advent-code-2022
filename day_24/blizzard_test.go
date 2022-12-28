package main

import (
	"testing"
)

func TestFindEndEasy(t *testing.T){
	input := "#.#####\n#.....#\n#>....#\n#.....#\n#...v.#\n#.....#\n#####.#"
	startingMap, blizzards := parseMap(input)
	cycle:= findCycleLength(startingMap)
	state := createTimeDimension(startingMap, cycle, blizzards)
	actual:= findLocation(Position{1, 0}, Position{len(startingMap[0])-2,len(startingMap)-1}, state, cycle, 0)
	want:=10
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted time to be %v`, input, actual, want )
	}
}
func TestFindEndHard(t *testing.T){
	input := "#.######\n#>>.<^<#\n#.<..<<#\n#>v.><>#\n#<^v^^>#\n######.#"
	startingMap, blizzards := parseMap(input)
	cycle:= findCycleLength(startingMap)
	state := createTimeDimension(startingMap, cycle, blizzards)
	actual:= findLocation(Position{1, 0}, Position{len(startingMap[0])-2,len(startingMap)-1}, state, cycle, 0)
	want:=18
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted time to be %v`, input, actual, want )
	}
}
func TestFindThereAndBackAndBackAgain(t *testing.T){
	input := "#.######\n#>>.<^<#\n#.<..<<#\n#>v.><>#\n#<^v^^>#\n######.#"
	startingMap, blizzards := parseMap(input)
	cycle:= findCycleLength(startingMap)
	state := createTimeDimension(startingMap, cycle, blizzards)
	start := Position{1, 0}
	end := Position{len(startingMap[0])-2,len(startingMap)-1}
	there:= findLocation(start, end, state, cycle, 0)
	back := findLocation(end, start, state, cycle, there)
	backAgain  := findLocation(start, end, state, cycle, back)

	want:=54
	if want != backAgain{
		t.Fatalf(`Given %v, got %v, but wanted time to be %v`, input, backAgain, want )
	}
}
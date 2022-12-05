package main

import (
	"testing"
	"fmt"
)

func TestUhg(t *testing.T) {
	var input string = `[N]    
[Z] [M]
 1   2

move 1 from 2 to 1`
	parsedStacks, parsedMoves := parseInput(input)
	//idek just print them
	fmt.Println(parsedStacks)
	fmt.Println(parsedMoves)
}

func TestFullParse(t *testing.T) {
	var input string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	parsedStacks, parsedMoves := parseInput(input)
	if(len(parsedMoves) != 4){
		t.Fatalf(`Didn't parse all moves correctly, given %v, got %v`, 4, parsedMoves)
	}
		n, z := parsedStacks[0].Pop(), parsedStacks[0].Pop()
	if n != "N" || z != "Z" {
		t.Fatalf(`First stack didnt parse correctly "%v" should b N, "%v" should be Z`, n, z)

	}
	d, c, m := parsedStacks[1].Pop(), parsedStacks[1].Pop(), parsedStacks[1].Pop()

	if d != "D" || c != "C" && m != "M" {
		t.Fatalf(`Second stack didnt parse correctly "%v%v%v"`, d, c, m)

	}
	p:= parsedStacks[2].Pop()
	if p != "P" {
		t.Fatalf(`Third stack didnt parse correctly "%v"`, p)

	}
}

func TestMoving(t *testing.T) {
	var stacks = []Stack{*New(), *New(), *New()}
	stacks[0].Push("Z")
	stacks[0].Push("N")
	stacks[1].Push("M")
	stacks[1].Push("C")
	stacks[1].Push("D")
	stacks[2].Push("P")
	var moves = []Reposition{Reposition{1,2,1}, Reposition{3,1,3}, Reposition{2,2,1}, Reposition{1,1,2}}
	want := "CMZ"
	executeRepositions(stacks, moves)
	actual:= stacks[0].Peek().(string) + stacks[1].Peek().(string) +stacks[2].Peek().(string)
	if want != actual{
		t.Fatalf(`Expected %v, got "%v"`, want, actual)
	}
}
func TestMovingNoOrderChange(t *testing.T) {
	var stacks = []Stack{*New(), *New(), *New()}
	stacks[0].Push("Z")
	stacks[0].Push("N")
	stacks[1].Push("M")
	stacks[1].Push("C")
	stacks[1].Push("D")
	stacks[2].Push("P")
	var moves = []Reposition{Reposition{1,2,1}, Reposition{3,1,3}, Reposition{2,2,1}, Reposition{1,1,2}}
	want := "MCD"
	executeRepositionsV2(stacks, moves)
	actual:= stacks[0].Peek().(string) + stacks[1].Peek().(string) +stacks[2].Peek().(string)
	if want != actual{
		t.Fatalf(`Expected %v, got "%v"`, want, actual)
	}
}
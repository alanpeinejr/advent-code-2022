package main

import (
	"testing" 
	"fmt"
)

func TestTailDoesntMoveOverlap(t *testing.T){
	input := [][]int{[]int{0}}
	inputTail:= Position{0,0}
	inputHead := Position{0,0}
	wantRope := [][]int{[]int{0}}
	wantTail := Position{0,0}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v`, input, inputTail, wantTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}
func TestTailDoesntMoveRightHorAdjacent(t *testing.T){
	input := [][]int{[]int{0, 0}}
	inputTail:= Position{0,0}
	inputHead := Position{0,1}
	wantRope := [][]int{[]int{0, 0}}
	wantTail := Position{0,0}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v`, input, inputTail, wantTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}

func TestTailDoesntMoveLeftHorAdjacent(t *testing.T){
	input := [][]int{[]int{0, 0}}
	inputTail:= Position{0,1}
	inputHead := Position{0, 0}
	wantRope := [][]int{[]int{0, 0}}
	wantTail := Position{0,1}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v`, input, inputTail, wantTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}

func TestTailDoesntMoveAboveVerAdjacent(t *testing.T){
	input := [][]int{[]int{0}, []int{0}}
	inputTail:= Position{1,0}
	inputHead := Position{0, 0}
	wantRope := [][]int{[]int{0}, []int{0}}
	wantTail := Position{1,0}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v`, input, inputTail, wantTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}
func TestTailDoesntMoveBelowVerAdjacent(t *testing.T){
	input := [][]int{[]int{0}, []int{0}}
	inputTail:= Position{0,0}
	inputHead := Position{1, 0}
	wantRope := [][]int{[]int{0}, []int{0}}
	wantTail := Position{0,0}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v`, input, inputTail, wantTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}
func TestTailDoesntMoveDiagonalAdjacent(t *testing.T){
	input := [][]int{[]int{0, 0}, []int{0, 0}}
	inputTail:= Position{0,0}
	inputHead := Position{1, 1}
	wantRope := [][]int{[]int{0, 0}, []int{0, 0}}
	wantTail := Position{0,0}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v, but is %v`, input, inputTail, wantTail, actualTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}
func TestTailDoesMoveUpAndLeft(t *testing.T){
	input := [][]int{[]int{0, 0}, []int{0, 0}, []int{0,0}}
	inputTail:= Position{1,2}
	inputHead := Position{0, 0}
	wantRope := [][]int{[]int{0, 0}, []int{1, 0}, []int{0,0}}
	wantTail := Position{0,1}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v, but is %v`, input, inputTail, wantTail, actualTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}
func TestTailDoesMoveDownAndRight(t *testing.T){
	input := [][]int{[]int{0, 0}, []int{0, 0}, []int{0,0}}
	inputTail:= Position{0,0}
	inputHead := Position{1, 2}
	wantRope := [][]int{[]int{0, 0}, []int{0, 1}, []int{0,0}}
	wantTail := Position{1,1}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v, but is %v`, input, inputTail, wantTail, actualTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}

func TestTailCountsAllVisits(t *testing.T){
	input := [][]int{[]int{0, 0}, []int{0, 1}, []int{0,0}}
	inputTail:= Position{0,0}
	inputHead := Position{1, 2}
	wantRope := [][]int{[]int{0, 0}, []int{0, 2}, []int{0,0}}
	wantTail := Position{1,1}
	actualRope, actualHead, actualTail := tailFollowsHead(input, inputHead, inputTail, true)
	if inputHead != actualHead {
		t.Fatalf(`Head should not change! Given %v, got %v`, inputHead, actualHead)
	}
	if wantTail != actualTail {
		t.Fatalf(`Input: %v, Given tail %v, should have moved to %v, but is %v`, input, inputTail, wantTail, actualTail)

	}
	for y, row:= range wantRope {
		for x, _ := range row {
			if wantRope[y][x] != actualRope[y][x]{
				t.Fatalf(`Rope didnt update correctly. Should be %v, was %v`, wantRope, actualRope)
			}
		} 
	}
}
func TestCountVisits(t *testing.T){
	input := "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2"
	rope:= [][]int{[]int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}}
	want:= 13
	actual:=countVisits(moveRope(input, rope, Position{0,4}, 1))
	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted %v`, input, actual, want)
	}
}
func TestMoveRope(t *testing.T){
	input := "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2"
	rope:= [][]int{[]int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}}
	// want:= [][]int{}
	actual := moveRope(input, rope, Position{0,4}, 1)
	fmt.Println(actual)
}
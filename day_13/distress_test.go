package main

import (
	"testing"
	"reflect"
)

func TestListBuildEmpty(t *testing.T){
	input := "[]"
	actual := buildList(input)
	want:= 0
	if want != len(actual) {
		t.Fatalf(`Given %v, got %v, but wanted length to be %v`, input, actual, want )
	}
}
func TestListBuildSingleNumber(t *testing.T){
	input := "[6]"
	actual := buildList(input)
	want:= 1
	if want != len(actual) && reflect.TypeOf(actual[0]).String() != "float64" {
		t.Fatalf(`Given %v, got %v, but wanted length to be %v`, input, actual, want )
	}
}
func TestListBuildNumberAndList(t *testing.T){
	input := "[6, []]"
	actual := buildList(input)
	want:= 2
	if want != len(actual) && reflect.TypeOf(actual[0]).String() != "float64" {
		t.Fatalf(`Given %v, got %v, but wanted length to be %v`, input, actual, want )
	}
}
func TestIsOrderedTwoEmptyLists(t *testing.T){
	left := "[]"
	right := "[]"
	actual := isOrdered(buildList(left), buildList(right))
	if actual > 0 {
		t.Fatalf(`Given left %v and right %v, got %v, but wanted %v`, left, right, actual, 0 )
	}
}
func TestIsOrderedTwoNumbersOrdered(t *testing.T){
	left := "[1]"
	right := "[2]"
	actual := isOrdered(buildList(left), buildList(right))
	if actual > 0 {
		t.Fatalf(`Given left %v and right %v, got %v, but wanted  %v`, left, right, actual, -1 )
	}
}
func TestIsOrderedTwoNumbersEqual(t *testing.T){
	left := "[2]"
	right := "[2]"
	actual := isOrdered(buildList(left), buildList(right))
	if actual != 0 {
		t.Fatalf(`Given left %v and right %v, got %v, but wanted %v`, left, right, actual, 0 )
	}
}
func TestIsOrderedTwoNumbersNotOrdered(t *testing.T){
	left := "[3]"
	right := "[2]"
	actual := isOrdered(buildList(left), buildList(right))
	if actual <= 0 {
		t.Fatalf(`Given left %v and right %v, got %v, but wanted %v`, left, right, actual, 1 )
	}
}

func TestIsOrderedOneListOneNumber(t *testing.T){
	left := "[[1]]"
	right := "[2]"
	actual := isOrdered(buildList(left), buildList(right))
	if actual > 0  {
		t.Fatalf(`Given left %v and right %v, got %v, but wanted %v`, left, right, actual, -1 )
	}
}
func TestIsOrderedOneListOneNumberUnordered(t *testing.T){
	left := "[[1, 2], 1]"
	right := "[1, 2]"
	actual := isOrdered(buildList(left), buildList(right))
	if actual <= 0 {
		t.Fatalf(`Given left %v and right %v, got %v, but wanted  %v`, left, right, actual, 1 )
	}
}

func TestOrdered2(t *testing.T){
	//[[1],[2,3,4]]
	// [[1],4]
	left := "[[1],[2,3,4]]"
	right := "[[1],4]"
	actual := isOrdered(buildList(left), buildList(right))
	if actual >= 0 {
		t.Fatalf(`Given left %v and right %v, got %v, but wanted %v`, left, right, actual, -1 )
	}
}


func TestSumming(t *testing.T){
	input := "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]"
	actual := orderedIndexes(input)
	want:= 13
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted length to be %v`, input, actual, want )
	}
}

func TestSortSignal(t *testing.T){
	input := "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]"
	actual := sortSignal(input)
	want:= 140
	if want != actual{
		t.Fatalf(`Given %v, got %v, but wanted length to be %v`, input, actual, want )
	}
}

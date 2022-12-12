package main

import (
	"testing"
)

func TestStartMiddleAllEdges(t *testing.T){
	input := "aaa\naSa\naaa"
	start := buildMap(input, 'S', assignIfEligible)
	want:= 1
	if want != start.Up.Value {
		t.Fatalf(`Given %v, got %v, but wanted Up to be %v`, input, start.Up.Value, want)
	}
	if want != start.Down.Value {
		t.Fatalf(`Given %v, got %v, but wanted Down to be %v`, input, start.Down.Value, want)
	}
	if want != start.Left.Value {
		t.Fatalf(`Given %v, got %v, but wanted Left to be %v`, input, start.Left.Value, want)
	}
	if want != start.Right.Value {
		t.Fatalf(`Given %v, got %v, but wanted Right to be %v`, input, start.Right.Value, want)
	}
}
func TestNilEdgeLeftTop(t *testing.T){
	input := "Saa\naaa\naaa"
	start := buildMap(input, 'S', assignIfEligible)
	want:= 1
	if nil != start.Up {
		t.Fatalf(`Given %v, got %v, but wanted Up to be %v`, input, start.Up, nil)
	}
	if want != start.Down.Value {
		t.Fatalf(`Given %v, got %v, but wanted Down to be %v`, input, start.Down.Value, want)
	}
	if nil != start.Left {
		t.Fatalf(`Given %v, got %v, but wanted Left to be %v`, input, start.Left, nil)
	}
	if want != start.Right.Value {
		t.Fatalf(`Given %v, got %v, but wanted Right to be %v`, input, start.Right.Value, want)
	}
}
func TestNilEdgeDownRight(t *testing.T){
	input := "aaa\naaa\naaS"
	start := buildMap(input, 'S', assignIfEligible)
	want:= 1
	if nil != start.Down {
		t.Fatalf(`Given %v, got %v, but wanted Down to be %v`, input, start.Down, nil)
	}
	if want != start.Up.Value {
		t.Fatalf(`Given %v, got %v, but wanted Up to be %v`, input, start.Up.Value, want)
	}
	if nil != start.Right {
		t.Fatalf(`Given %v, got %v, but wanted right to be %v`, input, start.Right, nil)
	}
	if want != start.Left.Value {
		t.Fatalf(`Given %v, got %v, but wanted left to be %v`, input, start.Left.Value, want)
	}
}

func TestPathLength(t *testing.T) {
	input := "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"
	start := buildMap(input, 'S', assignIfEligible)
	pathLength := pathLength(start.findEnd('E'))
	want:= 31
	if want != pathLength {
		t.Fatalf(`Given %v, got %v, but wanted  %v`, input, pathLength, want)
	}
}

func TestPathLengthReverseShort(t *testing.T) {
	input := "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"
	start := buildMap(input, 'E', assignIfEligibleReverse)
	pathLength := pathLength(start.findEnd('a'))
	want:= 29
	if want != pathLength {
		t.Fatalf(`Given %v, got %v, but wanted  %v`, input, pathLength, want)
	}
}
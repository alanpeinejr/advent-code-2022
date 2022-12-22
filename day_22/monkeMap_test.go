package main

import (
	"testing"
)

func TestSum(t *testing.T){
	input:= "        ...#    \n        .#..    \n        #...    \n        ....    \n...#.......#\n........#...\n..#....#....\n..........#.\n        ...#....\n        .....#..\n        .#......\n        ......#.\n\n10R5L5R10L4R5L5"
	me := pathFinder(parseLines(input))
	actual := score(me)
	want := 6032

	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted unmix to be %v`, input, actual, want )
	}

}
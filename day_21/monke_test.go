package main

import (
	"testing"
)

func TestSum(t *testing.T){
	input:= "root: pppw + sjmn\ndbpl: 5\ncczh: sllz + lgvd\nzczc: 2\nptdq: humn - dvpt\ndvpt: 3\nlfqf: 4\nhumn: 5\nljgn: 2\nsjmn: drzm * dbpl\nsllz: 4\npppw: cczh / lfqf\nlgvd: ljgn * ptdq\ndrzm: hmdt - zczc\nhmdt: 32"
	actual := moreMonke(parseLines(input))
	want := 152

	if want != actual {
		t.Fatalf(`Given %v, got %v, but wanted unmix to be %v`, input, actual, want )
	}

}
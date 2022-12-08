package main

import (
	"testing" 
	"fmt"
)

func TestBuildTreePatch(t *testing.T){
	var input string = "30373\n25512\n65332\n33549\n35390"
	want := [][]int{[]int{3,0,3,7,3}, []int{2,5,5,1,2}, []int{6,5,3,3,2}, []int{3,3,5,4,9}, []int{3,5,3,9,0}}
	actual := buildTreePatch(input)
	for i:=0; i < 4; i++{
		wantString := fmt.Sprintf("%v%v%v%v", want[i][0], want[i][1], want[i][2], want[i][3])
		actualString := fmt.Sprintf("%v%v%v%v", actual[i][0], actual[i][1], actual[i][2], actual[i][3])
			if wantString != actualString {
			t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
		}
	}

}

func TestIsNotVisible(t *testing.T){
	var input = "111\n101\n111"
	want := false
	actual := isVisible(buildTreePatch(input), Position{1,1})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
	
}

func TestIsLeftVisible(t *testing.T){
	var input = "010\n011\n010"
	want := true
	actual := isVisible(buildTreePatch(input),  Position{1,1})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestIsRightVisible(t *testing.T){
	var input = "010\n110\n010"
	want := true
	actual := isVisible(buildTreePatch(input),  Position{1,1})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestIsAboveVisible(t *testing.T){
	var input = "000\n111\n010"
	want := true
	actual := isVisible(buildTreePatch(input),  Position{1,1})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestIsBelowVisible(t *testing.T){
	var input = "010\n111\n000"
	want := true
	actual := isVisible(buildTreePatch(input),  Position{1,1})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestIsDoubleLeftNotVisibleButRightVisible(t *testing.T){
	//0010
	//0110
	//0010
	var input = "0010\n0110\n0010"
	want := true
	actual := isVisible(buildTreePatch(input),  Position{1,2})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestIsDoubleHorSidedNotVisible(t *testing.T){
	//0100
	//1010
	//0100
	var input = "0100\n1010\n0100"
	want := false
	actual := isVisible(buildTreePatch(input),  Position{1,1})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}

func TestIsDoubleVerSidedNotVisible(t *testing.T){
	//0100
	//0100
	//1010
	//0100
	var input = "0100\n0100\n1110\n0100"
	want := false
	actual := isVisible(buildTreePatch(input),  Position{1,2})
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
func TestIsCountVisible(t *testing.T){
	//30373
	//25512
	//65332
	//33549
	//35390
	var input string = "30373\n25512\n65332\n33549\n35390"
	want := 21
	actual := countVisibleTrees(buildTreePatch(input))
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}

func TestCountSeeableTreesAndScenicScore(t *testing.T){
	//30373
	//25512
	//65332
	//33549
	//35390
	var input string = "30373\n25512\n65332\n33549\n35390"
	var grid, tree  = buildTreePatch(input), Position{2, 3}
	treeHeight := grid[tree.y][tree.x]
	above := countSeeableTrees(reverseSlice(makeVerticleSlice(grid, 0, tree.y, tree.x)), treeHeight)
	below := countSeeableTrees(makeVerticleSlice(grid, tree.y + 1, len(grid), tree.x), treeHeight)
	left := countSeeableTrees(reverseSlice(grid[tree.y][0:tree.x]), treeHeight)
	right := countSeeableTrees(grid[tree.y][tree.x+1:len(grid[tree.y])], treeHeight)
	wantAbove := 2
	wantBelow := 1
	wantLeft := 2
	wantRight := 2
	if wantAbove != above {
		t.Fatalf(`Input above %v should return %v but was %v`, input, wantAbove, above)
	}
	if wantBelow != below {
		t.Fatalf(`Input below %v should return %v but was %v`, input, wantBelow, below)
	}
	if wantRight != right {
		t.Fatalf(`Input right %v should return %v but was %v`, input, wantRight, right)
	}
	if wantLeft != left {
		t.Fatalf(`Input left %v should return %v but was %v`, input, wantLeft, left)
	}
}

func TestCountSeeableForOne(t *testing.T){
	//30373
	//25512
	//65332
	//33549
	//35390
	var input string = "30373\n25512\n65332\n33549\n35390"
	var grid, tree  = buildTreePatch(input), Position{2, 3}
	actual := calculateScenicScore(grid, tree)
	want := 8
	if want != actual {
		t.Fatalf(`Input above %v should return %v but was %v`, input, want, actual)
	}
}
func TestIsScenicScore(t *testing.T){
	//30373
	//25512
	//65332
	//33549
	//35390
	var input string = "30373\n25512\n65332\n33549\n35390"
	want := 8
	actual := findHighesScenicScore(buildTreePatch(input))
	if want != actual {
		t.Fatalf(`Input %v should return %v but was %v`, input, want, actual)
	}
}
package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	//part 1, count the visible trees
	// fmt.Println(countVisibleTrees(buildTreePatch(readInput())))
	//part 2
	fmt.Println(findHighesScenicScore(buildTreePatch(readInput())))
	//268800 with 52 commented in, correct
	//4283136
}

func buildTreePatch(input string) [][]int{
	var patch [][]int = make([][]int, 0)
	rows := strings.Split(input, "\n")
	for _, rowString := range rows{
		row := make([]int, 0)
		for x:=0; x < len(rowString); x++{
			row = append(row, stringToInt(string(rowString[x])))
		}
		patch = append(patch, row)
	}
	return patch

}

func countVisibleTrees(patch [][]int) int {
	var totalVisible, xMax, yMax = 0, len(patch[0])-1, len(patch)-1
	for y, row := range patch {
		for x, _ := range row {
			if x == 0 || y == 0 || x == xMax || y == yMax {
				totalVisible+=1
			}else if isVisible(patch, Position{x, y}){
					totalVisible+=1
	
			}
		}
	}
	return totalVisible
}

func findHighesScenicScore(patch [][]int) int {
	var max =0
	for y, row := range patch {
		for x, _ := range row {
				treeScore := calculateScenicScore(patch, Position{x,y})
				fmt.Printf("%v\n", calculateScenicScore(patch, Position{x,y}))// why does observing it make it worK???
				if treeScore > max {
					max = treeScore
				}	
			}
		}
	return max
}

func calculateScenicScore(grid [][]int, tree Position) int {
	treeHeight := grid[tree.y][tree.x]
	above := makeVerticleSlice(grid, 0, tree.y, tree.x)
	below := makeVerticleSlice(grid, tree.y + 1, len(grid), tree.x)
	left := grid[tree.y][0:tree.x]
	right := grid[tree.y][tree.x+1:len(grid[tree.y])]
	//we want to reverse the left/above slices so we count from the tree and not the edge
	return countSeeableTrees(reverseSlice(left), treeHeight) *
	 countSeeableTrees(right, treeHeight) *
	 countSeeableTrees(reverseSlice(above), treeHeight) *
	 countSeeableTrees(below, treeHeight)

}

func countSeeableTrees(trees []int, treehouseHeight int) int {
	for index, tree := range trees{
		if tree >= treehouseHeight {
			return index + 1
		}
	}
	return len(trees)
}

func reverseSlice(s []int) []int{
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
	return s
}

func isVisible(grid [][]int, tree Position) bool{
	//test from above, below, left, and right
	treeHeight := grid[tree.y][tree.x]
	above := makeVerticleSlice(grid, 0, tree.y, tree.x)
	below := makeVerticleSlice(grid, tree.y + 1, len(grid), tree.x)
	left := grid[tree.y][0:tree.x]
	right := grid[tree.y][tree.x+1:len(grid[tree.y])]
	// fmt.Printf("left %v\n", isVisibleFrom(left, treeHeight))
	// fmt.Printf("right %v %v %v %v", isVisibleFrom(right, treeHeight), right, treeHeight, tree)
	// fmt.Printf("above %v\n", isVisibleFrom(above, treeHeight))
	// fmt.Printf("below %v\n", isVisibleFrom(below, treeHeight))
	// fmt.Println()


	return isVisibleFrom(left, treeHeight) || isVisibleFrom(right, treeHeight) || isVisibleFrom(above, treeHeight) || isVisibleFrom(below, treeHeight)
}

func isVisibleFrom(obstructors []int, tree int) bool{
	for _, obstruction := range obstructors {
		if obstruction >= tree {
			return false
		}
	}
	return true
}

func makeVerticleSlice(grid [][]int, from int, to int, column int) []int {
	slice := make([]int, 0)
	for y:=from; y < to; y++ {
		slice = append(slice, grid[y][column])
	}
	return slice
}



type (
	Position struct{
		x int
		y int
	}
)
//common helpers I'm copying because I haven't gotten around to fixing the local import problem
func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}
func readInput() string {
	var filename string
	if len(os.Args) < 2 {
        fmt.Println("Assuming local file input.txt")
		filename = "./input.txt"
    }else{
		filename = os.Args[1]
	}

    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Can't read file:", filename)
        panic(err)
    }

	//return and account for windows
	return strings.ReplaceAll(string(data), "\r\n", "\n")
}
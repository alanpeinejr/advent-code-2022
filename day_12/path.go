package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	//part 1
	// fmt.Println(pathLength(buildMap(readInput(), 'S',assignIfEligible).findEnd('E')))
	//part 2
	fmt.Println(pathLength(buildMap(readInput(), 'E', assignIfEligibleReverse).findEnd('a')))

}
//traverse the map like a 2d array creating a tree pointer for each element
	//fill out each tree's potential up/down/right/left 
	//locate S while doing this
//starting at S, BFS, return E on first encounter and trace it back to S counting the length
type (
	Tree struct {
    	Value  int
		Char   rune
		Up	   *Tree
		Down   *Tree
    	Right  *Tree
    	Left   *Tree
    	Parent *Tree
		Explored bool
	}
)
const heightString = "SabcdefghijklmnopqrstuvwxyzE"

func pathLength(end *Tree) int {
	node := end
	var length int
	for node.Parent != nil {
		length+=1
		node = node.Parent
	}
	return length
}

func (start *Tree) findEnd(endChar rune) (end *Tree) {
		queue := []*Tree{start}
		for len(queue) > 0 {
			current := queue[0]
			current.Explored = true
			if current != nil && current.Char == endChar {
				return current
			}
			queue = queue[1:]
				//and for each unexplored edge, set edges parent, queue
				if (current.Up != nil && !current.Up.Explored) {
					current.Up.Parent = current
					current.Up.Explored = true
					queue = append(queue, current.Up)
				}
				if (current.Down != nil && !current.Down.Explored) {
					current.Down.Parent = current
					current.Down.Explored = true
					queue = append(queue, current.Down)
				}
				if (current.Right != nil && !current.Right.Explored) {
					current.Right.Parent = current
					current.Right.Explored = true
					queue = append(queue, current.Right)
				}
				if (current.Left != nil && !current.Left.Explored) {
					current.Left.Parent = current
					current.Left.Explored = true
					queue = append(queue, current.Left)
				}
			
		}
		panic("Whoops, didnt find the end")
}

func buildMap(data string, startingRune rune, eligibleAssignment func(*Tree, *Tree)(*Tree)) *Tree {
	rows:= strings.Split(data, "\n")
	var start *Tree
	graph:= initGraph(len(rows[0]), len(rows))

	for y, row := range rows {
		for x, char := range row {
			//so we can assign down and right, we're gonna make those nodes prior
			node := graph[y][x]
			assignValue(node, char)

			//if its the start mark for return
			if char == startingRune {
				start = node
			}

			//now build edges
			if x == 0{
				node.Left = nil
			}else {
				node.Left = eligibleAssignment(node, graph[y][x-1])
			}
			if x == len(row) - 1 {
				node.Right = nil
			} else {
				assignValue(graph[y][x+1], rune(row[x + 1]))
				node.Right = eligibleAssignment(node, graph[y][x+1])
			}
			if y == 0 {
				node.Up = nil
			}else {
				node.Up = eligibleAssignment(node, graph[y-1][x])
			}
			if y == len(graph) - 1 {
				node.Down = nil
			}else {
				assignValue(graph[y+1][x], rune(rows[y+1][x]))
				node.Down = eligibleAssignment(node, graph[y+1][x])
			}
			graph[y][x] = node 
		}
	}
	return start
}

func initGraph(xBound int, yBound int) [][]*Tree {
	graph := make([][]*Tree, yBound)
	for y:= 0; y < yBound; y++ {
		graph[y] = make([]*Tree, xBound)
		for x := 0; x < xBound; x++ {
			graph[y][x] = new(Tree)
		}
	}
	return graph
}

func assignValue(node *Tree, char rune) {
	node.Char = char
	node.Value = strings.IndexRune(heightString, char)
	if char == 'S' {
		node.Value +=1
	} else if char == 'E'{
		node.Value -=1
	}
}

func assignIfEligible(node *Tree, nextNode *Tree) *Tree{
	if node.Value >= nextNode.Value - 1 {
		return nextNode
	}
	return nil	
}

func assignIfEligibleReverse(node *Tree, nextNode *Tree) *Tree{
	//because we can go down by 1, or back up by 1
	if nextNode.Value >= node.Value || nextNode.Value == node.Value - 1 {
		return nextNode
	}
	return nil	
}

func (tree *Tree) toString() string {
	return fmt.Sprintf(`Char: %v, Value: %v, Parent: %v, Up: %v, Down: %v, Right: %v, Left: %v`, tree.Char, tree.Value, tree.Parent, tree.Up, tree.Down, tree.Right, tree.Left)
}

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
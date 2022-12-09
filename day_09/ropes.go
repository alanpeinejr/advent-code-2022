package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	//part 1, count the visible trees
	// fmt.Println(countVisits(moveRope(readInput(), buildRope(1000), Position{500,500}, 1)))
	fmt.Println(countVisits(moveRope(readInput(), buildRope(1000), Position{500,500}, 9)))
}

func countVisits(rope [][]int) int {
	var count int
	for y:= 0; y < len(rope); y++{
		for x:=0; x < len(rope[y]); x++ {
			if rope[y][x] > 0 {
				count+=1
			}
		}
	}
	return count
}

func moveRope(input string,  rope [][]int, start Position, knots int) [][]int{
	var head, tails =start, make([]Position, knots)
	for i, _ :=range tails {
		tails[i] = start
	}
	rope[tails[0].y][tails[0].x]+=1//count the starting position	

	moveStrings := strings.Split(input, "\n")
	for _, moveString := range moveStrings {
		moveDirectionAndSize := strings.Split(moveString, " ")
		move:= moveDirectionAndSize[0]
		size := stringToInt(moveDirectionAndSize[1])
		for i:=0; i< size; i++{
			head = moveHead(move, head)
			//TODO modify this to operate on a slice of tails, only moveCount=true on last element
			//cascade, previous tail is new head
			for index, tail := range tails{
				moveCounts := index == len(tails) - 1 //the last tail only counts
				if index == 0 {
					rope, _, tails[index] = tailFollowsHead(rope, head, tail, moveCounts)
				}else {
					rope, _, tails[index] = tailFollowsHead(rope, tails[index-1], tail, moveCounts)

				}

			}
		}
	}
	return rope
}

func buildRope(size int) [][]int{
	//theres a cool world where I build it as it moves but this seems easier
	rope := make([][]int, size)
	for i:=0; i < size; i++ {
		rope[i] = make([]int, size)
	}
	return rope
}

func moveHead(direction string, head Position) Position{
	switch direction {
	case "L":
		return Position{ head.x-1, head.y}
	case "R":
		return Position{ head.x+1, head.y}
	case "U":
		return Position{ head.x, head.y-1}
	case "D":
		return Position{ head.x, head.y+1}
	default:
		return head
	}
}

func tailFollowsHead(rope [][]int, head Position, tail Position, countMove bool) ( [][]int,  Position,  Position) {
	var moved = false
	switch {
	//overlapping
	case head == tail:
		//do nothing
	//horizontally adjacent
	case intAbs(head.y - tail.y) <= 1 && intAbs(head.x - tail.x) <= 1:
		//adjacent, do nothing
	case head.y == tail.y && intAbs(head.x - tail.x) == 2:
		if(head.x > tail.x){ //move right
			tail.x+=1
		}else {				//move left
			tail.x-=1
		}
		moved = true
		//move left/right 1
	case head.x == tail.x && intAbs(head.y - tail.y) == 2:
		if( head.y > tail.y){//move down
			tail.y+=1
		} else{				//move up
			tail.y-=1
		}
		 moved = true
	//assuming diagonal
	default:
		if( head.y > tail.y){//move down
			tail.y+=1
		} else{				//move up
			tail.y-=1
		}
		if(head.x > tail.x){ //move right
			tail.x+=1
		}else {				//move left
			tail.x-=1
		}
		moved = true

	}
	if moved && countMove{
		rope[tail.y][tail.x]+=1
	}
	return rope, head, tail
}

type (
	Position struct{
		x int
		y int
	}
)
//
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

func intAbs(x int) int{
	if x < 0 {
		x = -x
	}
	return x
}
package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	// fmt.Println(score(pathFinder(parseLines(readInput()))))
	// fmt.Println(moreMonke(parseLines(readInput())))
	//part 2
	fmt.Println(score(pathFinder2E(parseLines(readInput()))))
	//8292 is not right, too low, had a wrong warp
	//192095 is not right, too hight, had a wrong warp
	//87281 is not right, too low, had a bug that allowed me to walk on rocks
	//104385 is correct

}

func score(me Me) int{
	return ((me.place.y+1) * 1000) + ( (me.place.x+1) * 4) + int(me.direction)
}

func pathFinder(monkeyMap [][]rune, instructions []string) Me {
	me := Me{right, findStart(monkeyMap)}

	for index, instruction := range instructions {
		if index%2 == 0 {
			//movement
			me = followInstruction(me, monkeyMap, stringToint(instruction))
		}else {
			//rotate
			me.direction = rotate(me.direction, instruction)
		}
	}

	return me
}
func pathFinder2E(monkeyMap [][]rune, instructions []string) Me {
	me := Me{right, findStart(monkeyMap)}

	for index, instruction := range instructions {
		if index%2 == 0 {
			//movement
			me = followInstructionCube(me, monkeyMap, stringToint(instruction))
		}else {
			//rotate
			me.direction = rotate(me.direction, instruction)
		}
	}

	return me
}

func rotate(facing Facing, direction string) Facing {
	switch direction {
	case "R":
		return clockwise[facing]
	case "L":
		return counterClockwise[facing]
	}
	panic("whoops")
} 

func findStart(monkeyMap [][]rune) Position {
	for y, row := range monkeyMap {
		for x, char := range row {
			if char == path {
				return Position{x,y}
			}
		}
	}
	//so it fails...an error would be better
	return Position{-1, -1}
}

func liveWithConsequences(me Me)  (Me) {
	//treat 50s as 3x4 where its found by x/50 and y/50 integer math
	//  (Y, X)
	//1 (0, 1) R&D as is, UP: 6, R; LEFT: 4, R
	//2 (0, 2) L as is, UP: 6, U; DOWN: 3, L; RIGHT 5, L
	//3 (1, 1) U&D as is, LEFT: 4 D ; Right: 2 U
	//4 (2, 0) R&D as is, UP: 3, R; Left: 1, R
	//5 (2, 1) U&L as is, DOWN: 6, L; RIGHT: 2, L
	//6 (3, 0) U as is, RIGHT: 5, U ; LEFT: 1, D ; DOWN 2, D
	x, y:= cubeMath(me.place)
	//first switch is "where are we?"
	//second is warp to where we're going
	switch {
	case y == 0 && x == 1:
		//1
		switch me.direction {
		case up:
			me = Me{right, Position{0, me.place.x - 50 + 150}}//this is redundant but helps me mentally
		case left://if were at y=0, need y =149 
			me = Me{right, Position{0, 149-me.place.y}}
		default:
			panic("whoops, wrong 1 bound")
		}
	case y == 0 && x == 2:
		//2
		switch me.direction {
		case up://if were at x=100, we need to be at x=0, was wrong, corrected
			me = Me{up, Position{me.place.x - 100, 199}}
		case down:
			me = Me{left, Position{99, me.place.x - 50}}
		case right:
			me = Me{left, Position{99, 149 - me.place.y}}
		default:
			panic("whoops, wrong 2 bound")
		}
	case y == 1 && x == 1:
		//3
		switch me.direction {
		case left://if were at 50, we need to be at 0, 
			me = Me{down, Position{me.place.y - 50, 100}}
		case right://if we're at y=50, we need to be at x=100, 
			me = Me{up, Position{me.place.y+50 , 49}}
		default:
			panic("whoops, wrong 3 bound")
		}
	case y == 2 && x == 0:
		//4
		switch me.direction {
		case up://if were at x = 0, we need to be at y = 50, 
			me = Me{right, Position{50, me.place.x + 50}}
		case left://if were at y=149, we need to be at y=0, y = 100 needs 49, 
			me = Me{right, Position{50,   149 - me.place.y}}
		default:
			 panic("woops, wrong 4 bound")
		}
	case y == 2 && x == 1:
		//5
		switch me.direction {
		case down://if wer at x = 50, we need to be at y = 150, , was wrong
			me = Me{left, Position{49, me.place.x + 100}}
		case right://if were at y=149, we need to be at y = 0, 
			me = Me{left, Position{149, 149 - me.place.y}}
		default:
			panic("whoops, wrong 5 bound")
		}
	case y == 3 && x == 0:
		//6
		switch me.direction {
		case right://if were at y= 199 we need x = 99, 
			me = Me{up, Position{me.place.y-100, 149}}
		case left:// y=199, x99, 
			me = Me{down, Position{me.place.y-100, 0}}
		case down://if we are at x = 0, we need to be at x=100
			me = Me{down, Position{me.place.x+100, 0}}
		default:
			panic("whoops, wrong 5 bound")
		}
	}

	return me
}

func cubeMath(position Position) (int, int) {
	return position.x/ 50, position.y /50
}

//haha, foresight. TODO: instead of find next path and warping there, return liveWithConsequences, which will
//recall this, with the new setup
//todo add a panic if we overbound to show we wrapped bad?
func followInstructionCube(me Me, monkeyMap [][]rune, instruction int) (Me) {
	if instruction == 0 {
		return me
	}
	var nextPosition Position
	var next, highbound int
	nextFace:= me.direction
	switch me.direction {
	case right:
		next = me.place.x + 1
		highbound = len(monkeyMap[me.place.y])
		nextPosition = Position{next, me.place.y}
	case left:
		next = me.place.x - 1
		highbound = len(monkeyMap[me.place.y])
		nextPosition = Position{next, me.place.y}
	case up:
		next = me.place.y - 1
		highbound = len(monkeyMap)
		nextPosition = Position{me.place.x, next}
	case down:
		next = me.place.y + 1
		highbound = len(monkeyMap)
		nextPosition = Position{me.place.x, next}
	}

	if next == highbound || next < 0 || monkeyMap[nextPosition.y][nextPosition.x] != rock && monkeyMap[nextPosition.y][nextPosition.x] != path {
		//wrap and find first char index
		warpedMe := liveWithConsequences(me)
		nextPosition = warpedMe.place
		nextFace = warpedMe.direction
	}
	//if next is a rock stop
	if monkeyMap[nextPosition.y][nextPosition.x] == rock {
		return me
	}

	return followInstructionCube(Me{nextFace, nextPosition}, monkeyMap, instruction-1)
}

//part 1
//this is kinda gross, would nodes have been better?...probably
func followInstruction(me Me, monkeyMap [][]rune, instruction int) (Me) {
	switch me.direction {
	case right:
		next := me.place.x + 1
		for x:=0; x < instruction; x++ {
			//figure out what next is if we need to wrap
			if next == len(monkeyMap[me.place.y]) || monkeyMap[me.place.y][next] != rock && monkeyMap[me.place.y][next] != path {
				//wrap and find first char index
				for i:=0; i < me.place.x; i++ {
					if monkeyMap[me.place.y][i] == rock || monkeyMap[me.place.y][i] == path {
						next = i
						break
					}
				}
			}
			//if next is a rock stop
			if monkeyMap[me.place.y][next] == rock {
				break
			}else {
				me.place = Position{next, me.place.y}
			}
			next+=1
		}
	case left:
		next := me.place.x - 1
		for x:=0; x < instruction; x++ {
			//figure out what next is if we need to wrap
			if next < 0  || monkeyMap[me.place.y][next] != rock && monkeyMap[me.place.y][next] != path {
				//wrap and find first char index
				for i:=len(monkeyMap[me.place.y]) - 1; i > me.place.x; i-- {
					if monkeyMap[me.place.y][i] == rock || monkeyMap[me.place.y][i] == path {
						next = i
						break
					}
				}
			}
			//if next is a rock stop
			if monkeyMap[me.place.y][next] == rock {
				break
			}else {
				me.place = Position{next, me.place.y}
			}
			next-=1
		}
	case up:
		next := me.place.y - 1
		for x:=0; x < instruction; x++ {
			//figure out what next is if we need to wrap
			if next < 0 || monkeyMap[next][me.place.x] != rock && monkeyMap[next][me.place.x] != path {
				//wrap and find first char index
				for i:=len(monkeyMap) - 1; i > me.place.y; i-- {
					if monkeyMap[i][me.place.x] == rock || monkeyMap[i][me.place.x] == path {
						next = i
						break
					}
				}
			}
			//if next is a rock stop
			if monkeyMap[next][me.place.x] == rock {
				break
			}else {
				me.place = Position{me.place.x, next}
			}
			next-=1
		}
	case down:
		next := me.place.y + 1
		for x:=0; x < instruction; x++ {
			//figure out what next is if we need to wrap
			if next == len(monkeyMap) || monkeyMap[next][me.place.x] != rock && monkeyMap[next][me.place.x] != path {
				//wrap and find first char index
				for i:=0; i < me.place.y; i++ {
					if monkeyMap[i][me.place.x] == rock || monkeyMap[i][me.place.x] == path {
						next = i
						break
					}
				}
			}
			//if next is a rock stop
			if monkeyMap[next][me.place.x] == rock {
				break
			}else {
				me.place = Position{me.place.x, next}
			}
			next+=1
		}
	}

	return me
}




func parseLines(input string)([][]rune, []string ){
	mapAndDirections := strings.Split(input, "\n\n")
	return parseMap(mapAndDirections[0]), parseDirections(mapAndDirections[1])

}
func parseMap(input string) [][]rune{
	var width, height int
	lines := strings.Split(input, "\n")
	height = len(lines)
	width = len(lines[0]) //my inputs first line is the longest
	monkeyMap := make([][]rune, height)
	for y, row := range lines {
		monkeyMap[y] = make([]rune, width)
		for x, char := range row {
			monkeyMap[y][x] = char
		}
	}
	return monkeyMap

}
func parseDirections(input string) []string {
	//will int convert in place, but this way its at least broken up
	input = strings.ReplaceAll(input, "L", " L ")
	input = strings.ReplaceAll(input, "R", " R ")
	return strings.Split(input, " ")
}
type ( 
	Facing	int
	Position struct{
		x int
		y int
	}
	Me		struct{
		direction 	Facing
		place		Position
	}
)
const (
	right      Facing = 0
	down       Facing = 1
	left       Facing = 2
	up         Facing = 3
	rock 		rune = '#'
	path 		rune = '.'
)
//lazy instancing for direction change
var clockwise  		 map[Facing]Facing = map[Facing]Facing{ 
	right: down,
	down: left,
	left: up,
	up: right,
}
var counterClockwise   map[Facing]Facing = map[Facing]Facing{ 
	right: up,
	down: right,
	left: down,
	up: left,
}

	
//common helpers I'm copying because I haven't gotten around to fixing the local import problem
func stringToint(this string) int {
	value, _ := strconv.Atoi(this)
	return int(value)
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
package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"bytes"
)

func main() {
	//part 2
	towerSize:= rocksFall(readInput(), 1000000, 7, 70000)
	fmt.Println(towerSize)
	//part 2
}

const line byte = byte('_')
const plus byte = byte('+')
const l byte = byte('L')
const pipe byte = byte('|')
const block byte = byte('o')
const stoppedRock byte = byte('#')
const left byte = byte('<')
const right byte = byte('>')
const empty byte = byte('.')

func rocksFall(input string, caveDepth int, caveWidth int, fallingRocks int) int {
	order := []byte{line, plus, l, pipe, block}
	cave := make([][]byte, caveDepth)
	fmt.Println(len(cave))
	for i:= 0; i < caveDepth; i++ {
		cave[i] = make([]byte, caveWidth)
	}
	rocksFallen := 0
	inputLength := len(input)
	inputCounter := 0
	highestRock := Position{0, len(cave)}//oob, but the floor is actually oob
	for i:= 0; i < fallingRocks; i++ {
		rock := drawRock(order[i % 5] , highestRock, cave)
		falling := true
		for falling { //fall until we stop
				falling, rock = moveRock(rock, cave, input[inputCounter % inputLength])
				rock = fall(rock, cave, falling)
		
			inputCounter+=1
		}
		highestRock = placeRock(rock, cave, highestRock.y)
		rocksFallen+=1
		if highestRock.y == 999099 || highestRock.y == 999099-2630 || highestRock.y == 999099 - 2*2630 {
			fmt.Println(rocksFallen)
		}
		if rocksFallen == 600 {
			fmt.Println(len(cave) - highestRock.y)
		}
		if rocksFallen == 600+1725+1000 {//left to 10t {
			fmt.Println(len(cave) - highestRock.y)

		}
		if rocksFallen == 600+1725+1725+1000 {//left to 10t {
			fmt.Println(len(cave) - highestRock.y)

		}
		if rocksFallen == 600+1725+1725+1725+1000 {//left to 10t {
			fmt.Println(len(cave) - highestRock.y)

		}
	}
		//thoughts on part 2
		//is there a pattern
	// output(cave)
	//noticed in my output
	//..####.
	//...#.#.
	//.#####.
	//.#.#.#. happens a lot, rythemically
	//narrowed it down to every 1725 a cycle starts
	//because of the floor throwing off the cycle, it doesnt start until 601
	//get hight at rocks fallen == 600
	//divide 10Trillion by 1725, then subrtract it from 10T to see how many bricks you need to reach that height
	//mine was 1600, but I knew hegiht of 600 bricks (900)
	//so get the height of my cycle after 1000 bricks (1525)
	//and add it to # cycles to 10T * my cycle height and wallah

	// fmt.Println(cave[len(cave) -10:len(cave)])
	//because we're actually falling down, len minus smallest y is our height
	fmt.Println(1000+600+(1000000000000/1725)*1725)
	fmt.Println(1525+ 900 + (1000000000000/1725) *2630)
	return len(cave) - highestRock.y
}

func moveRock(rock Rock, cave [][]byte, movement byte) (falling bool, newPosition Rock){
	//can we move the rock?
		//return stopped if a downward move and cant
	var dontMove bool
	falling = true
	newPosition.indexes = rock.indexes
	switch {
	case left == movement:
		for _, position := range rock.indexes {
			dontMove = position.x -1 < 0 || cave[position.y][position.x - 1] == stoppedRock
			if dontMove {
				break
			}
		}
		if !dontMove {
			for i, position := range rock.indexes {
				newPosition.indexes[i] = Position{position.x - 1, position.y}
			}
		}
	case right == movement:
		for _, position := range rock.indexes {
			dontMove = position.x + 1 > 6 || cave[position.y][position.x + 1] == stoppedRock
			if dontMove {
				break
			}
		}
		if !dontMove {
			for i, position := range rock.indexes {
				newPosition.indexes[i] = Position{position.x + 1, position.y}
			}
		}
	}
	for _, position := range newPosition.indexes {
		if  position.y + 1 >= len(cave) || cave[position.y + 1][position.x] == stoppedRock {
			falling = false
			break
		}
	}
	return
}
func fall(rock Rock, cave [][]byte, falling bool) (newPosition Rock){
	newPosition.indexes = rock.indexes
	if falling {
		for i, position := range newPosition.indexes {
			newPosition.indexes[i] = Position{position.x, position.y + 1}
		}

	}
	return

}

func placeRock(rock Rock, cave [][]byte, highestishPointish int) (newHighestPoint Position) {
	for _, position := range rock.indexes {
			cave[position.y][position.x] = stoppedRock
		}
		//rock could have moved lower than previous high, check it aroundish
		for y:=highestishPointish - 6; y < highestishPointish +10; y++ {
			for x:= 0; x < len(cave[y]); x++ {
				if cave[y][x] == stoppedRock {
					newHighestPoint = Position{x,y}
					return
				}
			}
		}
		return
}


func drawRock(rockToDraw byte, highestRock Position, cave [][]byte) (Rock)  {
	if(highestRock.y - 6 < 0) {
		panic("whoops, caves not deep enough")//probably gonna be a space problem part 2
	}
	var rock Rock
	switch  {
	//always two unites from left wall
	case rockToDraw == plus:
		rock = Rock{[]Position{Position{3, highestRock.y-4}, //bottom middle
							   Position{3, highestRock.y-5}, //middle middle
							   Position{3, highestRock.y-6}, //top middle
							   Position{4, highestRock.y-5}, //middle right
							   Position{2, highestRock.y-5}}} //middle left
	case rockToDraw == line:
		rock = Rock{[]Position{Position{2, highestRock.y-4}, //left most
							   Position{3, highestRock.y-4}, //middle left
							   Position{4, highestRock.y-4}, //middle right
							   Position{5, highestRock.y-4}}} //right most
	case rockToDraw == l:
		rock = Rock{[]Position{Position{2, highestRock.y-4}, //bottom left
							   Position{3, highestRock.y-4}, //bottom middle
							   Position{4, highestRock.y-4}, //bottom right
							   Position{4, highestRock.y-5}, //middle colum
							   Position{4, highestRock.y-6}}} //top column
	case rockToDraw == pipe:
		rock = Rock{[]Position{Position{2, highestRock.y-4}, //bottom 
							   Position{2, highestRock.y-5}, //bottom bottom middle
							   Position{2, highestRock.y-6}, //top top middle
							   Position{2, highestRock.y-7}}} //top
	case rockToDraw == block:
		rock = Rock{[]Position{Position{2, highestRock.y-4}, //bottom left
							   Position{3, highestRock.y-4}, //bottom right
							   Position{2, highestRock.y-5}, //top left
							   Position{3, highestRock.y-5}}} //top right				
	}
	return rock
}


type (
	Position struct {
		x int
		y int
	}
	Rock struct {
		indexes []Position
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
func output(cave [][]byte) {
	var null byte
	filename := "./output.txt"
	f, err := os.Create(filename)
    if err != nil {
        fmt.Println("Can't create file:", filename)
        panic(err)
    }
	for _, row := range cave {
		f.WriteString(fmt.Sprintf("%s\n", string(bytes.ReplaceAll(row, []byte{null}, []byte{empty}))) )
	}
	defer f.Close()
	//return and account for windows
}
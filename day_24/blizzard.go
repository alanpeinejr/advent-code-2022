package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	startingMap, blizzards := parseMap(readInput())
	cycle:= findCycleLength(startingMap)
	state := createTimeDimension(startingMap, cycle, blizzards)
	start := Position{1, 0}
	end := Position{len(startingMap[0])-2,len(startingMap)-1}
	minutes:= findLocation(start, end, state, cycle, 0)
	fmt.Println(minutes)
	//part 2
	minutesBack := findLocation(end, start, state, cycle, minutes)
	andBackAgain := findLocation(start, end, state, cycle, minutesBack)
	fmt.Println(andBackAgain)

}
	
//TODO E/W ore on a cycle l-2, N/S on a cycle h-2, so they repeat at lcm of L/H
//Calculate every possible state for cycle
//BFS where at this time can you move forward
//count explored, kill children that get consumed by the storm

func findLocation(start Position, end Position, blizzardOverTime [][][]mapIcon, cycle int, turn int) int {
	moves:= []Position{northMove, southMove, eastMove, westMove, wait}
	queue:= []PlaceInTime{{start, turn}}
	explored := map[PlaceInTime]bool{queue[0]:true}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.position.x == end.x && current.position.y == end.y {
			return current.time
		}

		//add children in all directions including wait if valid move
		for _, move := range moves {
			nextPlaceInTime := PlaceInTime{current.position.Add(move), current.time + 1}
			if !explored[nextPlaceInTime] && 
			   !nextPlaceInTime.position.outOfBounds(len(blizzardOverTime[current.time%cycle]), len(blizzardOverTime[current.time%cycle][0])) &&
			   blizzardOverTime[current.time%cycle][nextPlaceInTime.position.y][nextPlaceInTime.position.x] != wall &&
			   blizzardOverTime[(nextPlaceInTime.time)%cycle][nextPlaceInTime.position.y][nextPlaceInTime.position.x] == ground {
				queue = append(queue, nextPlaceInTime)
				explored[nextPlaceInTime] = true
			   }
		}

	}
	return -1
}

func createTimeDimension(blizzardMap [][]mapIcon, cycle int, blizzards []Blizzard) [][][]mapIcon{
	blizzardOverTime := make([][][]mapIcon, cycle)
	for t:=0; t < cycle; t++ {
		blizzardOverTime[t] = make([][]mapIcon, len(blizzardMap))
		for y:= 0; y < len(blizzardMap); y++{
			blizzardOverTime[t][y] = make([]mapIcon, len(blizzardMap[y]))
			for x:= 0; x < len(blizzardMap[y]); x++ {
				blizzardOverTime[t][y][x]= blizzardMap[y][x]
			}
		}
	}

	for t:= 0; t < cycle; t++ {
		//for each position, draw its position, then update it
		for i, blizzard := range blizzards {
			blizzardOverTime[t][blizzard.location.y][blizzard.location.x] = blizzard.direction
			//wrap logic and update blizzard location
			var heading Position
			switch blizzard.direction {
			case north:
				heading = northMove
			case south:
				heading = southMove
			case east:
				heading = eastMove
			case west:
				heading = westMove
			}

			nextPosition := wrapIfNeeded(blizzard.location.Add(heading), len(blizzardMap), len(blizzardMap[blizzard.location.y]))
			blizzards[i] = Blizzard{blizzard.direction, nextPosition}

		}

	}
	return blizzardOverTime
}

func wrapIfNeeded(position Position, height int, width int) Position{
	//a blizzard never enters the entrance/exit on my map
	//so can ignore the case of a blizzard needing to move into that before wrapping
	//-2 in new position to account for walls
	switch {
	case position.x == 0:
		return Position{width - 2, position.y}
	case position.x == width - 1:
		return Position{1, position.y}
	case position.y == 0:
		return Position{position.x, height-2}
	case position.y == height - 1:
		return Position{position.x, 1}
	default:
		return position
	}
}

func (this Position ) outOfBounds( height int, width int) bool {
	return this.x < 0 || this.x >= width || this.y < 0 || this.y >= height
}

func parseMap(input string) ([][]mapIcon, []Blizzard ){
	lines := strings.Split(input, "\n")
	blizzardMap := make([][]mapIcon, len(lines))
	blizzards:= make([]Blizzard, 0)
	for y, row := range lines {
		blizzardMap[y] = make([]mapIcon, len(lines[y]))
		for x, char := range row {
			switch  charToMapIcon[rune(char)]{
			case north:
				blizzards = append(blizzards, Blizzard{north, Position{x,y}})
				blizzardMap[y][x] = ground
			case south:
				blizzards = append(blizzards, Blizzard{south, Position{x,y}})
				blizzardMap[y][x] = ground
			case east:
				blizzards = append(blizzards, Blizzard{east, Position{x,y}})
				blizzardMap[y][x] = ground
			case west:
				blizzards = append(blizzards, Blizzard{west, Position{x,y}})
				blizzardMap[y][x] = ground
			default:
				blizzardMap[y][x] = charToMapIcon[rune(char)]
			}
		}

	}
	//fill in the ground

	return blizzardMap, blizzards
}

func findCycleLength(blizzardMap [][]mapIcon) int{
	//because of the border walls, -2
	//the map will always repeate after this length of time
	return lcm(len(blizzardMap) - 2, len(blizzardMap[0]) - 2)
}

//quick/lazy lcm
func lcm(one int, two int) int{
	lcm := 1
	if one > two {
		lcm = one
	} else {
		lcm = two
	}

	for {
		if lcm % one == 0 && lcm % two == 0 {
			return lcm
		}
		lcm++
	}
}

func (this Position) Add(that Position) Position {
	return Position{this.x+that.x, this.y + that.y}
}

type (
	mapIcon rune
	Blizzard struct {
		direction mapIcon
		location Position
	}
	PlaceInTime struct {
		position Position
		time int
	}
	Position struct {
		x	int
		y	int
	}
)
const (
	north mapIcon = '^'
	south mapIcon = 'v'
	east mapIcon = '>'
	west mapIcon = '<'
	ground mapIcon = '.'
	wall mapIcon = '#'
)

var charToMapIcon = map[rune]mapIcon{'^': north, 'v':south, '>':east, '<':west, '.':ground, '#':wall}
var (
	northMove       Position = Position{0,-1}
	southMove       Position = Position{0,1}
	eastMove        Position = Position{1, 0}
	westMove        Position = Position{-1,0}
	wait			Position = Position{0,0}
)
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
package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	// plot, elves:= tenRounds(parseMap(readInput(), 10))
	// minX, maxX, minY, maxY := findRectangle(elves)
	// fmt.Println(countGround(plot, minX, maxX, minY, maxY))
	//3970 correct
	//part 2
	fmt.Println(untilNoMoves(parseMap(readInput(), 100)))
	//923 first try, woot

}

func countGround(plot [][]rune, minX int, maxX int, minY int , maxY int) int{
	var count int
	var defaultRune rune
	for y:= minY; y <= maxY; y++ {
		for x:=minX; x <= maxX; x++{
			//check for default because i dont feel like prefilling with ground
			if plot[y][x] == ground  || plot[y][x] == defaultRune{
				count++
			}
		}
	}
	return count
}

func findRectangle(elves Elves) (int, int, int, int) {
	var minX, minY = 9223372036854775807, 9223372036854775807
	var maxX, maxY int
	for _, elf := range elves {
		if minX > elf.position.x {
			minX = elf.position.x
		}
		if maxX < elf.position.x {
			maxX = elf.position.x
		}
		if minY > elf.position.y {
			minY = elf.position.y
		}
		if maxY < elf.position.y {
			maxY = elf.position.y
		}
	}
	return minX, maxX, minY, maxY
}

func tenRounds(plot [][]rune, elves Elves) ([][]rune, Elves){
	order := []Heading{north, south, west, east}
	for i:=0; i< 10; i++{
		plot, elves, _ = round(plot, elves, order)
		//move the first element to the end for next round
		order = append(order[1:], order[0])
	}
	return plot, elves
}

func untilNoMoves(plot [][]rune, elves Elves) int {
	var noMoves bool
	roundNumber := 0
	order := []Heading{north, south, west, east}
	for true {
		plot, elves, noMoves = round(plot, elves, order)
		if(noMoves){
			return roundNumber + 1
		}
		//move the first element to the end for next round
		order = append(order[1:], order[0])
		roundNumber++
		if roundNumber %10 == 0 {
			//see if we're progressing
			fmt.Println(roundNumber)
		}
	}
	panic("we escaped an infinite loop")
}

func round(plot [][]rune, elves Elves, considerationOrder []Heading) ([][]rune, Elves, bool) {
	destinations := Destinations{}
	moves:=0
	//first half all the elves plan where to move
	for i, elf:= range elves {
		elves[i] = considerMoves(elf, plot, considerationOrder)
		destinations[elves[i].Add(elves[i].heading)]+=1
		if elves[i].heading != stationary {
			moves+=1
		}
	}

	//then they move, cannot move unless only elf wants to move there
	//Diplomacy rules
	for i, elf:= range elves {
		desitination := elf.Add(elf.heading)
		if destinations[desitination] == 1 {
			//elf moves
			plot[elf.position.y][elf.position.x] = ground
			elves[i].position = desitination
			plot[desitination.y][desitination.x] = elfMarker
		}
	}

	return plot, elves, moves == 0
}

func considerMoves(elf Elf, plot [][]rune, considerationOrder []Heading) Elf{
	validMoves := []Heading{}
	for _, heading := range considerationOrder {
		switch heading {
		case north:
			if considerHeading([]Position{elf.Add(north), elf.Add(north).Add(east), elf.Add(north).Add(west)}, plot) {
				validMoves = append(validMoves, north)
			}
		case south:
			if considerHeading([]Position{elf.Add(south), elf.Add(south).Add(east), elf.Add(south).Add(west)}, plot) {
				validMoves = append(validMoves, south)

			}
		case east:
			if considerHeading([]Position{elf.Add(east), elf.Add(east).Add(south), elf.Add(east).Add(north)}, plot) {
				validMoves = append(validMoves, east)

			}
		case west:
			if considerHeading([]Position{elf.Add(west), elf.Add(west).Add(south), elf.Add(west).Add(north)}, plot) {
				validMoves = append(validMoves, west)
			}
		}
	}
	var heading Heading
	//if we're surrounded, or not, we dont move
	if len(validMoves) == 4 || len(validMoves) == 0 {
		heading = stationary
	}else {
		heading = validMoves[0]
	}
	return Elf{elf.position, heading}
}

//returns true if there is no elves in any 3 spots
func considerHeading(headings []Position, plot [][]rune)  bool {
	for _, heading := range headings {
		if plot[heading.y][heading.x] == elfMarker {
			return false
		}
	}
	return true
}

func parseMap(input string, factor int) ([][]rune, Elves){
	var width, height int
	lines := strings.Split(input, "\n")
	elves := Elves{}
	height = len(lines)*factor
	width = len(lines[0])*factor
	plot := make([][]rune, height)
	//prefill every spot since some will be empty on all borders
	for y, _ := range plot {
		plot[y] = make([]rune, width)
	}
	startingX := len(lines) * factor/2 - len(lines)//middlesh offset by length right?
	startingY :=startingX
	for y, i:= startingY, 0; i < len(lines); y, i = y+1, i+1 {
		for x, j := startingX, 0; j < len(lines[i]); x, j = x+1, j+1{
			plot[y][x] = rune(lines[i][j])
			if rune(lines[i][j]) == elfMarker {
				elves = append(elves, Elf{Position{x,y}, stationary})
			}
		}
	}
	return plot, elves
}

type ( 
	Heading	Position
	Destinations map[Position]int
	Position struct{
		x int
		y int
	}
	Elf struct {
		position Position
		heading Heading
	}
	Elves []Elf
)

func (this Position) Add(that Heading) Position {
	return Position{this.x+that.x, this.y + that.y}
}
func (this Elf) Add(that Heading) Position {
	return Position{this.position.x+that.x, this.position.y + that.y}
}

var (
	north       Heading = Heading{0,-1}
	south       Heading = Heading{0,1}
	east        Heading = Heading{1, 0}
	west        Heading = Heading{-1,0}
	stationary	Heading = Heading{0,0}
	elfMarker 	rune = '#'
	ground 		rune = '.'
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
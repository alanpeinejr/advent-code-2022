package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	//part 2
	cave := buildCave(parseLines(readInput()))
	_, sandCount := landslide(cave, Position{500,0})
	fmt.Println(sandCount)
	output(cave)

}

func landslide(cave [][]rune, source Position) ([][]rune, int) {

	var defaultRune rune
	cave[source.y][source.x] = '+'
	var fallingSandPosition, atRestSand = source, 0
	for cave[source.y][source.x] != 'o' {
			//fall
			switch {
			case fallingSandPosition.y == len(cave) -1:
				//floor
				cave[fallingSandPosition.y][fallingSandPosition.x] = 'o'
				fallingSandPosition = source
				atRestSand+=1
			//fall below
			case cave[fallingSandPosition.y+1][fallingSandPosition.x] == defaultRune:
				fallingSandPosition = Position{fallingSandPosition.x, fallingSandPosition.y+1}
			//fall left
			case cave[fallingSandPosition.y+1][fallingSandPosition.x - 1] == defaultRune:
				fallingSandPosition = Position{fallingSandPosition.x - 1, fallingSandPosition.y+1}
			//fall right
			case cave[fallingSandPosition.y+1][fallingSandPosition.x + 1] == defaultRune:
				fallingSandPosition = Position{fallingSandPosition.x + 1, fallingSandPosition.y+1}
			default:
				cave[fallingSandPosition.y][fallingSandPosition.x] = 'o'
				fallingSandPosition = source
				atRestSand+=1
			}
		
	}
	return cave, atRestSand
}

func buildCave(rocks [][]Position, xMax int, yMax int) [][]rune {
	//fill the cave with air
	cave := make([][]rune, yMax)
	for y, _ := range cave {
		cave[y] = make([]rune, xMax)
	}
	//build rocks
	for y:= 0; y < len(rocks);  y++{
		//-1 because we're going to check forward
		for x:=0; x < len(rocks[y]) - 1; x++ {
			position := rocks[y][x]
			nextPosition := rocks[y][x+1]
			if position.x == nextPosition.x {
				//vertical line
				start := position.minY(nextPosition)
				end := position.maxY(nextPosition)
				for i:=start.y; i <= end.y; i++{
					cave[i][position.x] = '#'
				}
			} else if position.y == nextPosition.y {
				//horizontal line
				start := position.minX(nextPosition)
				end := position.maxX(nextPosition)
				for i:=start.x; i <= end.x; i++{
					cave[position.y][i] = '#'
				}
			}
		}

	}
	return cave
}

func (this Position) minX(that Position) Position {
	if this.x < that.x {
		return this
	}
	return that
}
func (this Position) maxX(that Position) Position {
	if this.x > that.x {
		return this
	}
	return that
}
func (this Position) minY(that Position) Position {
	if this.y < that.y {
		return this
	}
	return that
}
func (this Position) maxY(that Position) Position {
	if this.y > that.y {
		return this
	}
	return that
}

func parseLines(input string) ([][]Position, int, int ){
	var xMax, yMax int
	lineStrings := strings.Split(input, "\n")
	lines:= make([][]Position, len(lineStrings))

	for y, lineString := range lineStrings {
		vertexStringsArray := strings.Split(lineString, " -> ")
		lines[y] = make([]Position, len(vertexStringsArray))
		for x, vertexPairString := range vertexStringsArray {
			vertexPair := strings.Split(vertexPairString, ",")
			lines[y][x] = Position{stringToInt(vertexPair[0]), stringToInt(vertexPair[1])}
			//record the max's so we know our bounds
			if lines[y][x].x > xMax {
				xMax = lines[y][x].x
			}
			if lines[y][x].y > yMax {
				yMax = lines[y][x].y
			}
		}
	}
	//make sure we have a empty right side for potential stacking
	return lines, xMax + 200, yMax + 2
}

type (
	Position struct {
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

func output(cave [][]rune) {
	filename := "./output.txt"
	f, err := os.Create(filename)
    if err != nil {
        fmt.Println("Can't create file:", filename)
        panic(err)
    }
	for _, row := range cave {
		f.WriteString(fmt.Sprintf("%s\n", string(row)) )
	}
	//return and account for windows
}
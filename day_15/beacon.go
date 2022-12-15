package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main(){
	//part 1
	sensors, beacons:= parseLines(readInput())
	// cave := findImpossibleSpots(2000000, sensors, beacons)
	// fmt.Println(len(cave))
	//part 2
	 fmt.Println(findMissingBeacon(0, 4000000, sensors, beacons))

}

func findImpossibleSpots(row int, sensors []Position, beacons []Position) map[int]struct{} {
	set := make(map[int]struct{})
	for i, sensor := range sensors {
		// account for the neg x ind
		distance := sensor.calculateManhattenDistance(beacons[i])
		addToSetIfIncluded(set, row, sensor, distance)
		if(sensor.y == row) {
			delete(set, sensor.x)
		}
		if(beacons[i].y == row) {
			delete(set, beacons[i].x)
		}

	}
	return set
}

func (this Position) calculateManhattenDistance(that Position) int {
	return intAbs(this.x - that.x) + intAbs(this.y - that.y)
}
func addToSetIfIncluded(set map[int]struct{}, row int, position Position, radius int ){
	if (radius - intAbs(position.y - row)) * 2 + 1 > 0 { // calculates how many stars would be on the row +1 for center, *2 for both sides, abs(pos.y - the row ) to see how many levels away we are.
		//we've got values to add
		for x:=0 - (radius - intAbs(position.y - row)); x < radius - intAbs(position.y - row) + 1 ; x++ {
			drawX := position.x + x
			set[drawX] = struct{}{}
		}
	}
}
func findMissingBeacon(lowbound int, highbound int, sensors []Position, beacons[]Position) int{
	for y:= lowbound; y < highbound; y++ {
		ranges:= make(Range, 0)
		for i, sensor := range sensors {
			distance := sensor.calculateManhattenDistance(beacons[i])
			newRange := getRange(sensor, y, lowbound, highbound, distance)
			if len(newRange) > 0 {
				ranges  = append(ranges, newRange)
			}
			
		}
		isGap, gap := findGaps(ranges)
		if(isGap){
			return gap  * 4000000 + y
		}
		//every 1% so we can gauge progress
		if(y % 40000 == 0){
			fmt.Println(".")
		}
	}
	return 0
}
func findGaps(filled [][]int) (bool, int) {
	merged := merge(filled)
	if len(merged) > 1 {
		//we want the index after first element
		return true, merged[0][1] + 1

	}
	return false, -1

}

func getRange(position Position, y int, lowbound int, highbound int, radius int) []int{
	size:=(radius - intAbs(position.y - y)) * 2 + 1
	if size <= 0 {
		return []int{}
	}
	var xLow, xHigh int
	if position.x - size/2 < 0 {
		xLow = lowbound
	}else {
		xLow = position.x - size/2
	}
	if position.x + size/2 > highbound {
		xHigh = highbound
	}else {
		xHigh = position.x + size/2
	}
	return []int{xLow, xHigh}
}


func parseLines(input string) (sensors []Position, beacons []Position) {
	input = strings.ReplaceAll(input, "Sensor at ", "")
	input = strings.ReplaceAll(input, " closest beacon is at ", "")
	input = strings.ReplaceAll(input, "x=", "")
	input = strings.ReplaceAll(input, " y=", "")
	lines := strings.Split(input, "\n")
	sensors = make([]Position, len(lines))
	beacons = make([]Position, len(lines))
	for i, line := range lines {
		sensorBeaconStringArray := strings.Split(line, ":")
		sensorString := sensorBeaconStringArray[0]
		beaconString := sensorBeaconStringArray[1]
		sensorStringArray := strings.Split(sensorString, ",")
		sensors[i] = Position{stringToInt(sensorStringArray[0]), stringToInt(sensorStringArray[1])}
		beaconStringArray := strings.Split(beaconString, ",")
		beacons[i] = Position{stringToInt(beaconStringArray[0]), stringToInt(beaconStringArray[1])}
	}
	return 
}

type (
	Position struct {
		x 		int
		y 		int
	}
	Range [][]int
)
func (intA Range) Len() int {
	return len(intA)
}

func (intA Range) Swap(i, j int) {
	intA[i], intA[j] = intA[j], intA[i]
}

func (intA Range) Less(i, j int) bool {
	return intA[i][0] < intA[j][0]
}

func merge(intervals [][]int) [][]int {

	intA := Range(intervals)

	sort.Sort(intA)

	intervalsSorted := [][]int(intA)

	var output [][]int
	currentIntervalStart := intervalsSorted[0][0]
	currentIntervalEnd := intervalsSorted[0][1]
	for j := 1; j < len(intervalsSorted); j++ {
		if currentIntervalEnd >= intervalsSorted[j][0] {
			if intervalsSorted[j][1] > currentIntervalEnd {
				currentIntervalEnd = intervalsSorted[j][1]
			}
		} else {
			output = append(output, []int{currentIntervalStart, currentIntervalEnd})
			currentIntervalStart = intervalsSorted[j][0]
			currentIntervalEnd = intervalsSorted[j][1]
		}
	}
	output = append(output, []int{currentIntervalStart, currentIntervalEnd})
	return output

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

func intAbs(x int) int{
	if x < 0 {
		x = -x
	}
	return x
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
	defer f.Close()
}
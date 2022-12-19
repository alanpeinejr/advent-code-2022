package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
	"github.com/ernestosuarez/itertools"
)

func main() {
	rooms := parseLines(readInput())
	paths:= getAllShortestPaths(rooms)
	//because we'll never open the 0 rooms, we dont have to consider them
	roomsRemaining := make([]string, 0)
	for _, room := range rooms {
		if room.FlowRate != 0 {
			roomsRemaining = append(roomsRemaining, room.Name)
		}
	}
	//part 1
	// finalPressure := goToNextBest(rooms, paths, "AA", 30, 0, 0, 0, roomsRemaining)
	//part 2
	finalPressure := tooFatTooFurious(rooms, paths, 26, roomsRemaining)

	fmt.Println(finalPressure)




}
//im not sure if my surface can brute force this...
func tooFatTooFurious(rooms map[string]*Room, paths map[string]map[string]int, time int, closedValves []string) int{
	//max, we only have to go to half the *closed* (non zero) rooms, because the elephants will get the rest
	finalPressure := 0

	for i:= 1; i < len(closedValves)/2; i++ {
		for pathComb := range itertools.CombinationsStr(closedValves, i) {//every combination of length i for the rooms
			finalPressureMe := goToNextBest(rooms, paths, "AA", time, 0, 0, 0, pathComb)
			//so if we go to i length rooms, the elephant goes to all the other ones
			finalPressureElephant := goToNextBest(rooms, paths, "AA", time, 0, 0, 0, exclude(pathComb, closedValves))
			if finalPressureMe + finalPressureElephant > finalPressure {
				finalPressure = finalPressureMe + finalPressureElephant
			}

		}
	}
	return finalPressure
}

func exclude(this []string, from []string) []string{
	new := []string{}
	outer:
		for _, i := range from {
			for _, j := range this {
				if i == j {
					//i is in the excluded array, dont put it in
					continue outer
				}
			}
			new = append(new, i)
		}
	return new
}


func goToNextBest(rooms map[string]*Room, shortestPaths map[string]map[string]int, position string, timeLimit int, timePassed int, atTimePressure int, flowRate int, closedValves []string) int {
	currentScore := atTimePressure + (flowRate * (timeLimit - timePassed))
	maxScore := currentScore
	
	for _, roomName := range closedValves {
		travelTimePlusOpen := shortestPaths[position][roomName] + 1//time to open
		//dont look for rooms we cant get to
		if timePassed + travelTimePlusOpen < timeLimit  {
			postTravelTime := travelTimePlusOpen + timePassed
			//pressure increases while we walk there
			nextAtTimePressure := atTimePressure + travelTimePlusOpen * flowRate
			afterOpenFlow := flowRate + rooms[roomName].FlowRate
			nextScoreForRoom := goToNextBest(rooms, shortestPaths, roomName, timeLimit, postTravelTime, nextAtTimePressure, afterOpenFlow, deppend(closedValves, roomName))
			if (nextScoreForRoom > maxScore) {
				maxScore = nextScoreForRoom
			}

		}
	}

	return maxScore
}

func deppend(old []string, removed string) []string{
	copy := make([]string, 0)
	for _, copying := range old {
		if copying != removed {
			copy = append(copy, copying)
		}
	}
	return copy
}

//
func getAllShortestPaths(rooms map[string]*Room) (shortestPaths map[string]map[string]int) {
	//dont consider the 0 valve rooms as we'll never go to just them
	//add every node in the graph
	graph := dijkstra.NewGraph()
	var index int
	for _, room := range rooms {//i is a string when you range a map[string]
		room.Id = index //idk if range is same order every iterable, also saving in case i need later, graph lib needs int ids
		graph.AddVertex(room.Id)
		index ++
	}
	//add every connection in the graph
	for _, room := range rooms {
		for _, connection :=range room.Connections {
			graph.AddArc(room.Id, connection.Id, 1) //feels like edge could be weighted on inverse flow or something?
		}
	}

	shortestPaths = map[string]map[string]int{}

	//calculate the shortest path from every node, to every other node
	for nameA, roomA := range rooms {
		shortestPaths[nameA] = map[string]int{}
		fmt.Println(nameA)
		for nameB, roomB := range rooms {
			shortestpath, _ := graph.Shortest(roomA.Id, roomB.Id)
			shortestPaths[nameA][nameB] = int(shortestpath.Distance)
		}
	}


	return
}

func parseLines(input string) (rooms map[string]*Room) {
	input = strings.ReplaceAll(input, "Valve ", "")
	input = strings.ReplaceAll(input, " has flow rate", "")
	input = strings.ReplaceAll(input, " tunnels lead to valves ", "")
	input = strings.ReplaceAll(input, " tunnel leads to valve ", "")//why did they do two versions?

	//AA=0;AB,AC\n
	lines := strings.Split(input, "\n")
	rooms = map[string]*Room{}

	for _, line := range lines {
		roomsAndConnections := strings.Split(line, ";")
		roomAndFlowString := roomsAndConnections[0]
		connectionsString := roomsAndConnections[1]
		roomAndFlow := strings.Split(roomAndFlowString, "=")
		roomName := roomAndFlow[0]
		flow := roomAndFlow[1]
		connections:= strings.Split(connectionsString, ", ")


		//if we haven't seen the room yet
		if rooms[roomName] == nil {
			rooms[roomName] = new(Room)
		}
		room := rooms[roomName]
		room.Name = roomName
		room.FlowRate = stringToInt(flow)

		for _, connection := range connections {
			if rooms[connection] == nil {
				rooms[connection] = new(Room)
			}
			room.Connections = append(room.Connections, rooms[connection])
		}

	}
	return
}

type (
	Room struct {
		Name string
		FlowRate int
		Connections []*Room
		Id int
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


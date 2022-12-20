package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	fmt.Println(extractValue(parseLines(readInput()), 24))
	//part 2
	fmt.Println(extractValuer(parseLines(readInput())[:3], 32))

}
func extractValuer(blueprints []Blueprint, time int) int{
	product := 1
	for _, blueprint := range blueprints {
		product*= capitalism(blueprint, time)
	}
	return product
}

func extractValue(blueprints []Blueprint, time int) int{
	sum := 0
	for _, blueprint := range blueprints {
		sum+= qualityScore(blueprint.Id, capitalism(blueprint, time))
	}
	return sum
}

func qualityScore(id int, geoCollected int) int{
	return id * geoCollected
}

func capitalism(blueprint Blueprint, time int) int{
	//bfs where we queue all brancing potential turns, tracking the turn
	queue := []Turn{}
	//queue initial state
	queue = append(queue, Turn{
		Building : nop,
		Number : 0,
		Inventory : map[material]int{ore: 0, clay:0, obsidian:0, geode:0},
		Robits : map[material]int{ore:1, clay:0, obsidian:0, geode:0},
	})
	maxGeoCollected :=0
	//cache the turn so we can trim anwsered branches
	explored := make(map[cache]bool)

	for len(queue) > 0 {
		currentTurn := queue[0]
		queue = queue[1:]

		if currentTurn.Number > time {
			continue
		}

		if currentTurn.Inventory[geode] > maxGeoCollected {
			maxGeoCollected = currentTurn.Inventory[geode]
		}

		cacheHit := cache{
			Building : 	currentTurn.Building,
			Turn : 		currentTurn.Number,
			Orebit : 	currentTurn.Robits[ore],
			Claybit : 	currentTurn.Robits[clay],
			Obbit : 	currentTurn.Robits[obsidian],
			Geobit : 	currentTurn.Robits[geode],
		}

		if _, exists :=explored[cacheHit]; exists {
			continue
		}
		explored[cacheHit] = true

		//the workers produce
		currentTurn = produce(currentTurn)

		//the workers reproduce
		if currentTurn.Building != nop {
			currentTurn.Robits[currentTurn.Building]+=1
			currentTurn.Building = nop
		}

		potentialMoves:= getBuildOptions(currentTurn, blueprint)
		for _, potentialMove := range potentialMoves {
			nextTurn := Turn {
				Building : potentialMove,
				Number : currentTurn.Number + 1,
				Inventory : cloneMap(currentTurn.Inventory),
				Robits : cloneMap(currentTurn.Robits),
			}

			if potentialMove != nop {
				//consume
				for requiredResource, cost := range blueprint.Robits[potentialMove] {
					nextTurn.Inventory[requiredResource] -= cost
				}

			}
			queue = append(queue, nextTurn)

		}

	}
	return maxGeoCollected

}

func getBuildOptions(turn Turn, blueprint Blueprint) []material {
	options := []material{}
	outer:
	for _, resource := range order {
			robit := blueprint.Robits[resource]
			for requiredResource, cost := range robit {
				if turn.Inventory[requiredResource] < cost {
					continue outer
				}
			}
			options = append(options, resource)
	}
	//for when its better to save
	options =append(options, nop)
	return options

}

func produce(turn Turn) Turn{
	for _, resource := range order {
		turn.Inventory[resource] += turn.Robits[resource]
	}
	return turn
}

//generic clone i found
func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	nm := make(map[K]V)
	for k, v := range m {
		nm[k] = v
	}
	return nm
}

type (
	Robit map[material]int
	Blueprint struct {
		Id int
		Robits map[material]Robit
	}
	Turn struct {
		Building material
		Number int
		Inventory map[material]int
		Robits map[material]int
	}
	material string
	build material
	cache struct{
		Building material
		Turn int
		Orebit int
		Claybit int
		Obbit int
		Geobit int
	}
)
const (
	nop      material = "nop"
	ore      material = "ore"
	clay     material = "clay"
	obsidian material = "obsidian"
	geode    material = "geode"
)
var order []material = []material{
	geode,
	obsidian,
	clay,
	ore,
}
func (turn Turn) String() string {
	return fmt.Sprintf("turn: %d building: %s  robits: %v inv: %v", turn.Number, turn.Building, turn.Robits, turn.Inventory)
}
func parseLines(input string) []Blueprint {
	input = strings.ReplaceAll(input, "Blueprint ", "")
	input = strings.ReplaceAll(input, ": Each ore robot costs ", ",")
	input = strings.ReplaceAll(input, " ore. Each clay robot costs ", ",")
	input = strings.ReplaceAll(input, " ore. Each obsidian robot costs ", ",")
	input = strings.ReplaceAll(input, " ore and ", ",")
	input = strings.ReplaceAll(input, " clay. Each geode robot costs ", ",")
	input = strings.ReplaceAll(input, " obsidian.", "")
	//ID,Ore Robot Ore Cost,Clay Robot Ore Cost,Obsidian bot ore cost, Obsidian bot clay cost,geode bot ore cost, geode bot obsidian cost
	lines := strings.Split(input, "\n")
	blueprints := make([]Blueprint, len(lines))
	for i, line := range lines {
		blueprints[i].Robits = map[material]Robit{}
		
		idAndcostsArray := strings.Split(line, ",")
		id := stringToint(idAndcostsArray[0])
		oreBitOreCost := stringToint(idAndcostsArray[1])
		clayBitOreCost := stringToint(idAndcostsArray[2])
		obBitOreCost := stringToint(idAndcostsArray[3])
		obBitClayCost := stringToint(idAndcostsArray[4])
		geoBitOreCost := stringToint(idAndcostsArray[5])
		geoBitObCost := stringToint(idAndcostsArray[6])

		blueprints[i].Id = id
		blueprints[i].Robits[ore] = Robit{ore:oreBitOreCost}
		blueprints[i].Robits[clay] = Robit{ore:clayBitOreCost}
		blueprints[i].Robits[obsidian] = Robit{ore:obBitOreCost, clay: obBitClayCost}
		blueprints[i].Robits[geode] = Robit{ore:geoBitOreCost, obsidian: geoBitObCost}
	}
	return blueprints

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
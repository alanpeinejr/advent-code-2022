package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	// fmt.Println(moreMonke(parseLines(readInput())))
	//part 2
	lazyHuman(parseLines(readInput()))

}

func lazyHuman(available AvailableNumbers, monkies WaitingMonkeys) {
	//0
	i :=3375719472000
	for true {
		ogAvailable := cloneMap(available)
		ogAvailable["humn"] = i
		moreMonke(ogAvailable, monkies)
		i+=1
		//i+=1000000000000

	}
	//3375719472770
}

func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	nm := make(map[K]V)
	for k, v := range m {
		nm[k] = v
	}
	return nm
}

func moreMonke(available AvailableNumbers, monkies WaitingMonkeys) int {
	if root, exits := available["root"]; exits{
		//we found root
		return root
	}
	lessMonk := []Monkey{}
	for _, monke := range monkies {
		left, leftExists := available[monke.Left]
		right, rightExists := available[monke.Right]
		if leftExists && rightExists {
			if monke.Name == "root"{
				fmt.Println(available[monke.Left], available[monke.Right], available["humn"])
				if available[monke.Left] <= available[monke.Right] {
					panic(available["humn"])
				}

			}
			available[monke.Name] = monkeyBusines(monke, left, right)
		}else {
			lessMonk = append(lessMonk, monke)
		}
	}
	//got thru the list, call again until root
	return moreMonke(available, lessMonk)

}

func monkeyBusines(monkey Monkey, left int, right int) int {
	switch {
	case monkey.Operation == "+":
		return left + right
	case monkey.Operation == "-":
		return left - right
	case monkey.Operation == "*":
		return left * right
	case monkey.Operation == "/":
		return left / right
	default:
		panic("whoops")
	}
}

func parseLines(input string)(AvailableNumbers, WaitingMonkeys){
	lines := strings.Split(input, "\n")
	waiting := WaitingMonkeys{}
	available := AvailableNumbers{}
	//hcvr: jvhm * sflb
	//vclr: 8
	for _, line := range lines {
		monkeyAndFormula := strings.Split(line, ": ")
		monkeyName := monkeyAndFormula[0]
		if len(line) > 10 {
			//must be waiting
			operation := string(monkeyAndFormula[1][5])
			leftAndRight := strings.Split(monkeyAndFormula[1], " " + operation+ " ")
			waiting = append(waiting,  Monkey{monkeyName, operation, leftAndRight[0], leftAndRight[1]})
		}else {
			available[monkeyName] = stringToint(monkeyAndFormula[1])
		}
	}
	return available, waiting
}

type (
	AvailableNumbers map[string]int
	WaitingMonkeys []Monkey
	Monkey struct {
		Name string
		Operation string
		Left string
		Right string
	}
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
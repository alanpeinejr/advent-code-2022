package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	//part 1
	fmt.Println(calculateCombinedSignalStrengths(executeInstructions(readInput()), []int{20, 60, 100, 140, 180, 220 } ))
	render(executeInstructions(readInput()))
}

func calculateCombinedSignalStrengths(registerValues []int, intervals []int) int{
	var count int
	for _, interval :=range intervals {
		count+= registerValues[interval-1] * interval
	}
	return count
}

func render(registerValues []int){
	var cycles int
	for y:=0; y<6; y++{
		for x:=0; x < 40; x++ {
			
			if registerValues[cycles] == x || registerValues[cycles] == x -1 || registerValues[cycles] == x + 1 {
				fmt.Printf("#")
			}else {
				fmt.Printf(".")
			}

			cycles+=1
		}
		fmt.Printf("\n")
	}

}

func executeInstructions(input string) []int{
	instructions := strings.Split(input, "\n")
	registerValues := make([]int, 0)

	register := 1

	for _, instruction := range instructions {
		cyclesTaken, amountToUpdate := cycle(register, instruction)
		if cyclesTaken == 1 {
			registerValues = append(registerValues, register)
		}else if cyclesTaken ==2 {
			registerValues = append(registerValues, register, register)
		}
		register = updateRegister(register, amountToUpdate)
	}
	return registerValues
}

func cycle(register int, instruction string) (int, int){
	if instruction == "noop"{
		return 1, 0
	}else{
		addAndAmountArray := strings.Split(instruction, " ")
		amount := stringToInt(addAndAmountArray[1])
		return 2, amount
	}
}

func updateRegister(register int, update int) int {
	return register + update
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
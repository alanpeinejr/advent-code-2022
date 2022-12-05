package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	// fmt.Println(peekTops(executeRepositions(parseInput(readInput()))))
	fmt.Println(peekTops(executeRepositionsV2(parseInput(readInput()))))
}

type Reposition struct {
	quantity int
	from int
	to int
}

func executeRepositions(stacks []Stack, moves []Reposition) []Stack {
	for _, move := range moves {
		for i := 0; i < move.quantity; i++{
			stacks[move.to -1].Push(stacks[move.from - 1].Pop())
		}
	}
	return stacks
}

func executeRepositionsV2(stacks []Stack, moves []Reposition) []Stack {
	for _, move := range moves {
		var craneGrasp [] interface{}
		for i := 0; i < move.quantity; i++{
			craneGrasp = append(craneGrasp, stacks[move.from - 1].Pop())
			// stacks[move.to -1].Push(stacks[move.from - 1].Pop())
		}
		for j := len(craneGrasp) -1; j >= 0; j--{
			stacks[move.to - 1].Push(craneGrasp[j])
		}
	}
	return stacks
}

func peekTops(stacks []Stack) string {
	var output string
	for _, stack := range stacks {
		output += stack.Peek().(string)
	}
	return output
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

func parseInput(input string) ([]Stack, []Reposition){
	var manipulation string
	//remove "move "
	manipulation = strings.ReplaceAll(input, "move ", "")
	//remove " from "
	manipulation = strings.ReplaceAll(manipulation, " from ", ",")
	//remove " to "
	manipulation = strings.ReplaceAll(manipulation, " to ", ",")
	//remove brackets
	// manipulation = strings.ReplaceAll(manipulation, "[", " ")
	// manipulation = strings.ReplaceAll(manipulation, "]", " ")

	//seperate stacks from the moves
	stacksAndMoves := strings.Split(manipulation, "\n\n")
	stackString := stacksAndMoves[0]
	moveStringBlob := stacksAndMoves[1]
	//seperate stacks into lines
	stackLevelStrings := strings.Split(stackString, "\n")
	//from the last line, load stacks
	outputStacks := make([]Stack, 0)
	for i:=len(stackLevelStrings)-1; i >= 0; i--{
		//first iteration is just the amount of stacks
		if(i == len(stackLevelStrings) - 1) {
			stackStringArray := strings.Split(stackLevelStrings[i], "   ")
			for j := 0; j < len(stackStringArray); j++ {
				outputStacks = append(outputStacks, *New())
			}
		}else{
			//get the chars in positions "[M] [N]""
									   //"-1---2""
			var chars int
			for j:= 1; j < len(stackLevelStrings[i]); j+=4{
				crate := strings.TrimSpace(string(stackLevelStrings[i][j]))
				if len(crate) > 0{
					// fmt.Println(crate, j)
					outputStacks[chars].Push(crate)
				}
				chars +=1;
			}
		}

	}

	//parse moves
	movesStringArray := strings.Split(moveStringBlob, "\n")
	movesOutput := make([]Reposition, len(movesStringArray))
	for i, moveString := range movesStringArray {
		moveStringArray:= strings.Split(moveString, ",")
		//you can just make a 2d array in go? wtf
		move := Reposition{
			quantity:stringToInt(moveStringArray[0]),
		 	from:stringToInt(moveStringArray[1]),
		 	to:stringToInt(moveStringArray[2])}
		movesOutput[i] = move
	}
	return outputStacks, movesOutput
}

func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}

//something is massively wrong with my go dev environment and it wont allow me to import an external package, nor an internal one

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}	
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}
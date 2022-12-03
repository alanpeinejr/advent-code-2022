package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(rummageRucksack(readInput()))
}
//priority is lowercase first, annoyingly opposite of asci, 0 becaue score starts at 1
const scoreString = "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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

func rummageRucksack(data string) (score int) {
	sacks := strings.Split(data, "\n")
	scorer := getPriority()
	for _, sack := range sacks {
		score = scorer(findDuplicateType(sack))
	}
	return
}

func findDuplicateType(rucksack string) string {
	compartment1 := rucksack[:len(rucksack)/2]
	compartment2 := rucksack[len(rucksack)/2:]

	duplicate := compartment1[strings.IndexAny(compartment1, compartment2)];
	return string(duplicate)
}

func getPriority() func(typeChar string) int {
	score := 0
	return func(typeChar string) int {
		indexScore := strings.IndexAny(scoreString, typeChar)
		if(indexScore == -1){
			panic("whoops")
		}
		score+= indexScore
		return score
	}
}
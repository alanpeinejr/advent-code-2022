package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//part 1
	//fmt.Println(rummageRucksack(readInput()))
	//part 2
	fmt.Println(findBadges(readInput()))
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

func findBadgeType(sacks []string) string{
	if(3 != len(sacks)){
		panic("whoops, not 3 sacks")
	}
	hashmap := make(map[rune]int)
	for index, sack := range sacks {
		for _, letter := range sack {
			if(index == len(sacks) - 1 && hashmap[letter] == 2){
				return string(letter)
			}
			//only record the first instance of it per sack
			if(hashmap[letter] == index){
				hashmap[letter] +=1
			}
		}
	}
	//if we got here we didnt find a common badge
	panic("whoops, no badge")
}

func findBadges(input string) (score int) {
	sacks := strings.Split(input, "\n")
	scorer := getPriority()
	for index:=0; index+3 <= len(sacks); index+=3 {
		score = scorer(findBadgeType(sacks[index:index+3]))
	}
	return
}
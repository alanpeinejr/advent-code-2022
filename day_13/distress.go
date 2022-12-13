package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"sort"
)

func main(){
	//part 1
	fmt.Println(orderedIndexes(readInput()))
	//part 2
	fmt.Println(sortSignal(readInput()))
}

func sortSignal(input string) int {
	noDoubleLines:=  strings.ReplaceAll(input, "\n\n", "\n")
	noDividers:= strings.Split(noDoubleLines, "\n")
	fullList:= append(noDividers, "[[2]]", "[[6]]")

	sort.Slice(fullList, func(i, j int) bool {
		return isOrdered(buildList(fullList[i]), buildList(fullList[j])) < 0
	})

	var two, six int
	for i, item :=range fullList {
		if item == "[[2]]" {
			two = i + 1
		}
		if item == "[[6]]" {
			six = i + 1
		}
	}

	return two * six


}

func orderedIndexes(input string) int {
	pairs := strings.Split(input, "\n\n")
	var sum int
	for i, pair := range pairs {
		pairArray := strings.Split(pair, "\n")
		if isOrdered(buildList(pairArray[0]), buildList(pairArray[1])) <= 0{
			sum += i + 1//because index is 1 based
		}
	}
	return sum

}

//Single line
func buildList(input string) []any {
	var data []any
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func isOrdered(leftI interface{}, rightI interface{}) int{
	left, leftIsFloat := leftI.(float64); 
	right, rightIsFloat := rightI.(float64);
	switch {
	case leftIsFloat && rightIsFloat:
		if left < right {
			return -1
		}else if right < left {
			return 1
		}
		return 0
	case leftIsFloat && !rightIsFloat:
		return isOrdered([]interface{}{left}, rightI)
	case !leftIsFloat && rightIsFloat:
		return isOrdered(leftI, []interface{}{right})
	default:
		var ordered int
		//both items are lists
		leftList, _ := leftI.([]interface{})
		rightList, _ := rightI.([]interface{})
		for i, _ := range leftList {
			if i > len(rightList) - 1 {
				//if right ends before left
				return 1
			}
			//compound the list checks, ordered and equal are equivalent here
			ordered = isOrdered(leftList[i], rightList[i])

			if(ordered < 0) {
				return -1
			}
			//if our recursive list call gave us any false, we can just end it early
			if ordered > 0 {
				return 1
			}
		}
		if len(leftList) < len(rightList){
			//list left ended sooner, correct order
			return -1
		} 

		//no decision made at this level
		return 0
	}
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

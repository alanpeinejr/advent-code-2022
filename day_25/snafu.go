package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	fmt.Println(sum(parseInput(readInput())))
	//part 2

}

func sum(values []string) string {
	//for every number, right align them, sum as a string (in place?) from left most
	carryOver:="0"
	result:= ""
	for _, value := range values {
		if len(value) > len(result) {
			result = padRight(result, len(value) - len(result))
		}else if len(result) > len(value) {
			value = padRight(value, len(result) - len(value))
		} else {
			//same length, give them space to carry
			value = padRight(value, 1)
			result = padRight(result, 1)

		}
		var nextResult = []rune(result)
		for j:= len(value) -1; j >= 0; j--{ 
			nextResult[j], carryOver = add(string(result[j]), string(value[j]), carryOver)
		}
		result = string(nextResult)
		carryOver="0"
	}
	return result
}

func add(a string, b string, c string) (rune, string) {
	sum := fmt.Sprintf("%v", stringToint(snafuNumbers[a].value) + stringToint(snafuNumbers[b].value) + stringToint(snafuNumbers[c].value))
	//theres surely a cleaner way to do this string/rune/byte int stuff
	return rune(snafuNumbers[sum].value[0]), snafuNumbers[sum].carryover 
}
func padRight(numberString string, amount int) string {
	//fmt.Sprintf("%0*s", amount, numberString) doesnt do what I want it to, every time?
	for i:=0; i < amount; i++{
		numberString = "0" + numberString
	}
	return numberString
}

//it'll be a pain to convert this into base 10, add it all, then convert back to this negative having base 5...
//or I could just add it up in base 5
func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

type (
	Additive struct {//strings because i think? it will make me code the addition better and not mess up using regular numbers 
		value string
		carryover string
	}
)

//the highest carryover possible is 1/- account for the sums when carryover is required
var snafuNumbers = map[string]Additive{
	"5": Additive{"0", "1"},
	"4": Additive{"-1", "1"},
	"3": Additive{"=", "1"},
	"2": Additive{"2", "0"},
	"1": Additive{"1", "0"},
	"0": Additive{"0", "0"},
	"-": Additive{"-1", "0"},
	"-1": Additive{"-", "0"},//for sum of -1 to output correctly
	"=": Additive{"-2", "0"},
	"-2": Additive{"=", "0"},//for sum of -2
	"-3": Additive{"2", "-"},
	"-4": Additive{"1", "-"},
	"-5": Additive{"0", "-"},
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
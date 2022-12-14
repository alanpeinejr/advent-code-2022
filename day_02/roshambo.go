package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(simulateMatches(readInput()))
}
var scoreMap = map[string]int{
	"A":     1,
	"B":     2,
	"C":     3,
	"X":     1,
	"Y":     2,
	"Z":     3,
}

var loseMap = map[string]string{
	"A":	"Z",
	"B":	"X",
	"C":	"Y",
}
var winMap = map[string]string{
	"A":	"Y",
	"B":	"Z",
	"C":	"X",
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

func simulateMatches(data string) (score int) {
	matches := strings.Split(data, "\n")
	scorer := matchWinScore()
	for _, match := range matches {
		opponentsHand := string(match[0])
		//lose, draw, win
		outcome := string(match[2]);
		var myThrow string;
		switch outcome{
		case "X":
			myThrow = loseMap[opponentsHand]
		case "Y":
			myThrow = opponentsHand
		case "Z": 
			myThrow = winMap[opponentsHand]
		}
		score = scorer(scoreMap[opponentsHand], scoreMap[myThrow])
	}

	return
}

func matchWinScore() func(opponentsThrow int, yourThrow int) int {
	score:= 0


	return func(opponentsThrow int, yourThrow int) int{
		//how much was our throw worth?
		score+= yourThrow

		//did we win?
		switch {
		//tie
		case opponentsThrow == yourThrow:
			score+=3
		//you throw Rock to win
		case yourThrow == scoreMap["X"] && opponentsThrow == scoreMap["C"]:
			score+=6
		//you throw Paper to win
		case yourThrow == scoreMap["Y"] && opponentsThrow == scoreMap["A"]:
			score +=6
		//you throw Scissors to win
		case yourThrow == scoreMap["Z"] && opponentsThrow == scoreMap["B"]:
			score+=6
		default:
		//the rest are losing cases and dont mattter
		}

		return score;
	}
}
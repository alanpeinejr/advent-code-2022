package main

import (
	"fmt"
)

func main() {
	//part 1
	monkes:= buildMonke()
	fmt.Println(calculateMonkeyBusiness(monkes))
	
}

func calculateMonkeyBusiness(monkes []*Monke) int{
	for i:=0; i < 20; i++ {//part 1 its 20, 2 its 10000
		for _, monke := range monkes {
			returnToMonke(monke, monkes)
		} 
	}
	return monkeyBusiness(monkes)
}

func returnToMonke(monke *Monke, monkes []*Monke) {
	monke.itemsInspected += len(monke.items)
	for _, item := range monke.items {
		worryAssessment := monke.operation(item)
		worryAssessment = worryAssessment/3 //part 2 no longer reduces
		var itemDestIndex int
		if(monke.test(worryAssessment)) {
			itemDestIndex = monke.positiveThrow
		} else {
			itemDestIndex = monke.negativeThrow
		}
		monkes[itemDestIndex].items = append(monkes[itemDestIndex].items, worryAssessment)

	}
	//they always throw all items
	monke.items = make([]int, 0)
}

func monkeyBusiness(monkes []*Monke) int {
	var first, second int
	for _, monke := range monkes {
		monkeyBusiness := monke.itemsInspected
		switch {
		case first < monkeyBusiness:
			second = first
			first = monkeyBusiness
		case second > monkeyBusiness:
			second = monkeyBusiness
		default:
			//do nothings
		}
		fmt.Println(*monke)
	}
	return first * second
}

//data strucutre for the problem, lazy human entered way
type (
	Monke struct {
		items []int
		operation func( int) int
		test func(int) bool 
		positiveThrow int
		negativeThrow int
		itemsInspected int
	}
)

func buildMonke() []*Monke {
	monkes := []*Monke{
		&Monke{ //0
			items: []int{56, 52, 58, 96, 70, 75, 72},
			operation: func(input int) int {
				return input * 17},
			test: func(input int) bool {
				return input % 11 == 0},
			positiveThrow: 2,
			negativeThrow: 3,
		},
		&Monke{ //1
			items: []int{75, 58, 86, 80, 55, 81},
			operation: func(input int) int {
				return input + 7},
			test: func(input int) bool {
				return input % 3 == 0},
			positiveThrow: 6,
			negativeThrow: 5,
		},
		&Monke{ //2
			items: []int{73, 68, 73, 90},
			operation: func(input int) int {
				return input * input},
			test: func(input int) bool {
				return input % 5 == 0},
			positiveThrow: 1,
			negativeThrow: 7,
		},
		&Monke{ //3
			items: []int{72, 89, 55, 51, 59},
			operation: func(input int) int {
				return input + 1},
			test: func(input int) bool {
				return input % 7 == 0},
			positiveThrow: 2,
			negativeThrow: 7,
		},
		&Monke{ //4
			items: []int{76, 76, 91},
			operation: func(input int) int {
				return input * 3},
			test: func(input int) bool {
				return input % 19 == 0},
			positiveThrow: 0,
			negativeThrow: 3,
		},
		&Monke{ //5
			items: []int{88},
			operation: func(input int) int {
				return input + 4},
			test: func(input int) bool {
				return input % 2 == 0},
			positiveThrow: 6,
			negativeThrow: 4,
		},
		&Monke{ //6
			items: []int{64, 63, 56, 50, 77, 55, 55, 86},
			operation: func(input int) int {
				return input + 8},
			test: func(input int) bool {
				return input % 13 == 0},
			positiveThrow: 4,
			negativeThrow: 0,
		},
		&Monke{ //7
			items: []int{79, 58},
			operation: func(input int) int {
				return input + 6},
			test: func(input int) bool {
				return input % 17 == 0},
			positiveThrow: 1,
			negativeThrow: 5,
		},
	}
	return monkes

}

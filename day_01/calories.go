package main

import ( 
	"fmt"
 	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(findMaxCalories(read()))
}

func read() string {
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


func findMaxCalories(inventory string) int {
	//split according to all the blank lines
	var maxCalories, elvesHolding = 0, strings.Split(inventory, "\n\n")
	for _, elf := range elvesHolding {
		elfsBag := strings.Split(elf, "\n")
		elfsSum := sum(elfsBag)
		if( maxCalories < elfsSum) {
			maxCalories = elfsSum
		}

	}
	return maxCalories;
}

func sum(values []string) int {
	var sum int
	for _, value := range values {
		var i, err = strconv.Atoi(strings.TrimSpace(value))
		if(err != nil){
			fmt.Println(`Whoops, $s`, err)
		}
		sum+= i
	}
	return sum
}
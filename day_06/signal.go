package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//part 1
	// fmt.Println(findPacketIdentifierEndIndex(readInput()))
	//part 2
	fmt.Println(findMessageIdentifierEndIndex(readInput()))

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

func findPacketIdentifierEndIndex(message string) int{
	for i:= 4; i <= len(message); i++{
		if !hasDuplicates(message[i-4:i]){
			return i
		}
	}
	return -1
}
func findMessageIdentifierEndIndex(message string) int{
	for i:= 14; i <= len(message); i++{
		if !hasDuplicates(message[i-14:i]){
			return i
		}
	}
	return -1
}

func hasDuplicates(marker string) bool {
	for index, char := range marker {
		//if this char index isnt the lest index, theres a duplicate
		if index != strings.LastIndex(marker, string(char)){
			return true
		}
	}
	return false
}
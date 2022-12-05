package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	//part 1
	// fmt.Println(countFullContainingSections(readInput()))
	//part 2
	fmt.Println(countOverlappingSections(readInput()))
}
type SectionRange struct{
	lowerBound int
	upperBound int
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

func countFullContainingSections(data string) int {
	var count int
	pairs := strings.Split(data, "\n")
	for _, pair := range pairs {
		section1, section2 := getSectionRanges(pair)
		if(section1.contains(section2) || section2.contains(section1)){
			count+=1
		}
	}
	return count
}
func countOverlappingSections(data string) int {
	var count int
	pairs := strings.Split(data, "\n")
	for _, pair := range pairs {
		section1, section2 := getSectionRanges(pair)
		if(section1.overlaps(section2)){
			count+=1
		}
	}
	return count
}

func getSectionRanges(pairString string) ( SectionRange,  SectionRange){
	//32-33,34-36
	pairStringArray := strings.Split(pairString, ",")
	section1Array := strings.Split(pairStringArray[0], "-")
	section2Array := strings.Split(pairStringArray[1], "-")
	section1lower, _ := strconv.Atoi(section1Array[0])
	section1upper, _ := strconv.Atoi(section1Array[1])
	section2lower, _ := strconv.Atoi(section2Array[0])
	section2upper, _ := strconv.Atoi(section2Array[1])
	return SectionRange{section1lower, section1upper}, SectionRange{section2lower, section2upper}
}

func (outerSection SectionRange) contains(innerSection SectionRange) bool {
	return outerSection.lowerBound <= innerSection.lowerBound && outerSection.upperBound >= innerSection.upperBound
}

func (outerSection SectionRange) overlaps(innerSection SectionRange) bool {
	return outerSection.containsInt(innerSection.lowerBound) ||
	outerSection.containsInt(innerSection.upperBound) ||
	innerSection.containsInt(outerSection.lowerBound) ||
	innerSection.containsInt(outerSection.upperBound)
	
}

func (section SectionRange) containsInt(value int) bool {
	return section.lowerBound <= value && section.upperBound >= value 
}
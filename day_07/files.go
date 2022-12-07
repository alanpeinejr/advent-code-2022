package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"math"
)

func main(){
	//part 1
	files := buildFileSystem(readInput())
	files.calculateSize()
	// fmt.Println(findSizeOfDirectoriesLessThanTenHundredThousand(files))
	fmt.Println(findSmallestDirForNeededSpace(70000000, 30000000, files).size)
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

func findSmallestDirForNeededSpace(maxSpace int, neededSpace int, files *File) *File {
	minToDelete := neededSpace - (maxSpace - files.size)
	if(minToDelete < 0){
		panic("whoops, probaly not right")
	}


	return findMinDirLargerThan(minToDelete, files)
}

func findMinDirLargerThan(value int, files *File) *File {
	var returnValue = files
	if len(files.sub) == 0 {
		return New("notMe", math.MaxInt)
	}
	for _, file := range files.sub {
		var subMin = findMinDirLargerThan(value, file)
		if subMin.size < returnValue.size  && subMin.size > value {
			returnValue = subMin
		}
	}
	return returnValue
}


func findSizeOfDirectoriesLessThanTenHundredThousand(fileSystem *File) int {
	var size int
	if len(fileSystem.sub) == 0 {
		return 0
	}
	//count self?
	if fileSystem.size < 100000 {
		size+= fileSystem.size
	}

	for _, file := range fileSystem.sub {
			//count the sub directories
			size += findSizeOfDirectoriesLessThanTenHundredThousand(file)
	}
	return size

	 
}

func buildFileSystem(explorationCommands string) *File{
	//for every line
		//if it starts with $ handle action
		//otherwise, we make a new file, add it to the current File
	var fileSystem *File = New("/", 0)
	var currentDir *File 
	commands := strings.Split(explorationCommands, "\n")
	for _, command := range commands{
		switch {
		case strings.HasPrefix(command, "$ cd /"):
			currentDir = fileSystem
		case strings.HasPrefix(command, "$ cd .."):
			currentDir = currentDir.parent
		case strings.HasPrefix(command, "$ cd ")://if we're here must be a file nav
			currentDir = currentDir.goTo(command[5:])
		case strings.HasPrefix(command, "$ ls"):
			//haha why
		case strings.HasPrefix(command, "dir "):
			fileName:= command[4:]
			currentDir.addFile(New(fileName, 0))
		default:
			//should only get here if its a regular file
			sizeAndName := strings.Split(command, " ")
			currentDir.addFile(New(sizeAndName[1], stringToInt(sizeAndName[0])))
		}
	}
	return fileSystem
}


type (
	File struct{
		sub []*File
		size int
		name string
		parent *File
	}
)

func New(fileName string, size int) *File{
	file:= new(File)
	file.sub = make([]*File, 0)
	file.name = fileName
	file.size = size
	return file
}

func (this *File) addFile(that *File){
	that.parent = this
	this.sub = append(this.sub, that)
}

func (this File) goTo(name string) *File {
	for _, destination := range this.sub {
		if name == destination.name {
			return destination
		}
	}
	panic("Whoops, no such file")
}

func (this *File) calculateSize() int{
	if(len(this.sub) == 0){
		return this.size
	}
	for _, file := range this.sub {
		 this.size+= file.calculateSize()
	}
	return this.size
}
func stringToInt(this string) int {
	value, _ := strconv.Atoi(this)
	return value
}

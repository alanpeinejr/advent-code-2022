package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	list, zero := createNodes(parseLines(readInput()))
	setNeighbors(list)
	// unmixAll(list)
	// fmt.Println(sum(zero))
	//part 2
	decrypt(list, 811589153)
	superUnmixAll(list)
	fmt.Println(sum(zero))

}

func superUnmixAll(list []*Node){
	for i := 0; i < 10; i++ {
		unmixAll(list)
	}
}

func decrypt(list []*Node, key int) []*Node{
	for _, number := range list {
		number.value *= key
	}
	return list
}

func move(number *Node, movement int) *Node {
	for movement < 0 {
		number = number.prev
		movement++
	}
	for movement > 0 {
		number = number.next
		movement--
	}
	return number
}

func unmix(number *Node, l int) {
	//remove ourselves
	temp := number.prev
	number.prev.next = number.next
	number.next.prev = number.prev

	//go fwd/back
	temp = move(temp, number.value%(l-1))//because we've exluded temp

	//put in place
	number.prev = temp
	number.next = temp.next
	number.prev.next = number
	number.next.prev = number
}

func unmixAll(list []*Node) {
	for _, number := range list {
		unmix(number, len(list))
	}
}

func setNeighbors(list []*Node) {
	list[0].prev = list[len(list)-1]
	list[len(list)-1].next = list[0]
	for i := 1; i < len(list); i++ {
		list[i].prev = list[i-1]
		list[i-1].next = list[i]
	}
}

func sum(n0 *Node) int {
	sum := 0
	for i, n := 0, n0; i < 3; i++ {
		n = move(n, 1000)
		sum += n.value
	}
	return sum
}

func createNodes(data []int) ([]*Node, *Node){
	var order []*Node
	var zero *Node
	for _, value := range data {
		node := &Node{value: value }
		if node.value == 0 {
			zero = node
		}
		order = append(order, node)
	}
	return order, zero
}

type (
	Node struct {
		value     	int
		prev 		*Node
		next 		*Node
	}
)

func parseLines(input string)[]int{
	lines := strings.Split(input, "\n")
	output := make([]int, len(lines))
	for i, line := range lines {
		output[i] = stringToint(line)
	}
	return output
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
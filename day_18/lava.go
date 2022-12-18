package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	//part 1
	positions:= getPositions(readInput())
	blob := makeBlob(positions)
	sides:= countSides(blob, positions)
	fmt.Println(sides)
	//part 2
	fmt.Println(countSidesNoInner(blob, positions))

}

func getPositions(input string) ([]Position) {
	lines := strings.Split(input, "\n")
	positions := make([]Position, len(lines))
	for i, line:= range lines {
		values := strings.Split(line, ",")
		positions[i] = Position{stringToint(values[0]), stringToint(values[1]), stringToint(values[2])}
	}
	return positions
}

func makeBlob(positions []Position) [][][]int{
	//find bounds
	var maxX, maxY, maxZ int
	for i, _:= range positions {
		if(maxX < positions[i].x) {
			maxX = positions[i].x
		}
		if(maxY < positions[i].y) {
			maxY = positions[i].y
		}
		if(maxZ < positions[i].z) {
			maxZ = positions[i].z
		}
	}
	//give edge room
	maxX+=2
	maxY+=2
	maxZ+=2
	//make blob
	blob:= make([][][]int, maxX)
	for i, _ := range blob {
		blob[i] = make([][]int, maxY)
		for j, _ := range blob[i] {
			blob[i][j] = make([]int, maxZ)
		}
	}

	for _, position := range positions {
		blob[position.x][position.y][position.z]=1
	}

	return blob
}

func countSidesNoInner(blob [][][]int, positions []Position) (sides int) {
	//count all the sides
	//for every air, check if its 'within'
	sides = countSides(blob, positions)
	checkedAir := []Position{}
	pockets := [][]Position{}
	for x:=0; x < len(blob); x++ {
		for y:= 0; y < len(blob[x]); y++ {
			for z:= 0; z < len(blob[x][y]); z++ {
				if blob[x][y][z] == 0 {
					//found air
						//check if its in my found pockets
					position := Position{int(x),int(y),int(z)}
					 if isWithin(position, checkedAir) {
						//we found it in previous hole check that found a border
						continue
					 }
					if isWithinAny(position, pockets) {
						//we found it, its already in a hole
						continue
					}
					foundHole, hole := bfsHole(position, blob)
					if foundHole{
						pockets = append(pockets, hole)

					}else {
						//dont need to recheck the positions again because they lead to a border
						checkedAir = append(checkedAir, hole...)
					}
					

				}

			}
		}
	}
	for _, pocket := range pockets {
		blob:= makeBlob(pocket)
		sides -=  countSides(blob, pocket)
	}

return

}
func isWithin(position Position, pocket []Position) bool{
	for _, spot :=range pocket {
		if spot.x == position.x && spot.y == position.y && spot.z == position.z {
			return true
		}
	}
	return false
}

func isWithinAny(position Position, pockets [][]Position) bool{
	for _, pocket := range pockets {
		if isWithin(position, pocket) {
			return true
		}
	}
	return false
}

func bfsHole(origin Position, blob [][][]int) (bool, []Position) {
	// potentialHole := []Position{origin}
	queue := []Position{origin}
	explored := []Position{origin}

	//if we can find a border there isns a hole
	for len(queue) > 0 {
		current := queue[0]
		if foundBorder(current, blob){
			//if we hit a border its not a hole
			return false, explored
		}
		queue = queue[1:]
		//add this children if they aren't explored, if we hit a border, cant be a hole
		child := Position{current.x + 1, current.y, current.z}
		if !isWithin(child, explored)  && child.x  < int(len(blob)) && blob[child.x][child.y][child.z] != 1 {
			queue = append(queue, child)
			explored = append(explored, child)
		}
		child = Position{current.x - 1, current.y, current.z}
		if !isWithin(child, explored) && child.x > 0 && blob[child.x][child.y][child.z] != 1 {
			queue = append(queue, child)
			explored = append(explored, child)
		}
		child = Position{current.x, current.y + 1, current.z}
		if !isWithin(child, explored) && child.y < int(len(blob[child.x]))&& blob[child.x][child.y][child.z] != 1{
			queue = append(queue, child)
			explored = append(explored, child)
		}
		child = Position{current.x, current.y - 1, current.z}
		if !isWithin(child, explored) && child.y >  0 && blob[child.x][child.y][child.z] != 1{
			queue = append(queue, child)
			explored = append(explored, child)
		}
		child = Position{current.x, current.y, current.z+1}
		if !isWithin(child, explored) &&  child.z < int(len(blob[child.x][child.y])) && blob[child.x][child.y][child.z] != 1{
			queue = append(queue, child)
			explored = append(explored, child)
		}
		child = Position{current.x, current.y, current.z-1}
		if !isWithin(child, explored) && child.z > 0 && blob[child.x][child.y][child.z] != 1{
			queue = append(queue, child)
			explored = append(explored, child)
		}
	}
	//how many hole spots we found
	return true, explored

}

func foundBorder(position Position, blob [][][]int) bool{
	if position.x <= 0 || position.x >= int(len(blob) -1){
		return true
	}
	if position.y <= 0 || position.y >= int(len(blob[position.x]) -1) {
		return true
	}
	if position.z <= 0 || position.z >= int(len(blob[position.x][position.y]) -1) {
		return true
	}

	return false
}
func makePocket(xmin int, xmax int, ymin int, ymax int, zmin int, zmax int) []Position {
	position := []Position{}
	for x:=xmin+1; x < xmax; x++ {
		for y:= ymin+1; y < xmax; y++ {
			for z:= zmin+1; z < xmax; z++ {
				position = append(position, Position{int(x), int(y), int(z)})
			}
		}
	}

	return position


}
func countSides(blob [][][]int, positions []Position) (sides int) {
	for _, position := range positions {
		if(blob[position.x+1][position.y][position.z] != 1) {
			sides+=1
		}	
		if(position.x ==0 || blob[position.x-1][position.y][position.z] != 1) {
			sides+=1
		}
		if(blob[position.x][position.y+1][position.z] != 1) {
			sides+=1
		}	
		if(position.y == 0 || blob[position.x][position.y-1][position.z] != 1) {
			sides+=1
		}	
		if(blob[position.x][position.y][position.z+1] != 1) {
			sides+=1
		}	
		if(position.z == 0 || blob[position.x][position.y][position.z-1] != 1) {
			sides+=1
		}	
	}

	return 
}

type (
	Position struct {
		x int
		y int
		z int
	}
)
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
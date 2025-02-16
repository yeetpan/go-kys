package main

import (
	"fmt"
)

func main() {
	var fruitList = []string{"a", "c", "banana"}
	//fmt.Printf("Type of fruitlist is %T", fruitList)

	fruitList = append(fruitList, "Mango", "peach")
	fmt.Println(fruitList)
	// python kind of slicing.
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	// rather ususal syntax for declaration
	highScores := make([]int, 4)
	highScores[0] = 1
	highScores[1] = 12
	highScores[2] = 13
	highScores[3] = 41
	//(throws errors)highScores[4]=199

	// below method re-assigns
	highScores = append(highScores, 555, 666)
	// fmt.Println(highScores)

	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	//remove value from slices based on index.

	var courses = []string{"r", "js", "swift", "py", "ruby"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

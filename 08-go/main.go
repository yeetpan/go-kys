package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "a"
	fruitList[1] = "p"
	fruitList[3] = "e"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list is", len(fruitList))

	var vegList = [3]string{"p", "n", "e"}
	fmt.Println(len(vegList))
}

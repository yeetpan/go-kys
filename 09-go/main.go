package main

import "fmt"

// maps in goLang
func main() {
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["RB"] = "Ruby"

	fmt.Println(languages["TS"])

	//delete
	delete(languages, "RB")
	fmt.Println(languages)

	//loop through maps
	// for key, value := range languages {
	// 	fmt.Printf("For Key %v,value is %v\n", key, value)
	// }

	for _, value := range languages {
		fmt.Printf("Value is %v", value)
	}
}

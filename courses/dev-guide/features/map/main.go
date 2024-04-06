// Map == Hash (Ruby) == Object (js) == Dict (Python)
package main

import "fmt"

//Comparison against structs
/*
Map - All keys must be same type | Struct - Values can be of different type
Map - DO not need to know all keys at compile time | Struct - need to know different fields at compile time
Map - Reference Type | Struct - Value Type
Map - represent collection of related properties | Struct - represent "thing" with alot of different properties
Map - keys are indexed (can iterate) | Struct - keys do not support indexing

Others:
Map - All values must be of same type.

*/

func main() {
	//Ways to declare a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	//Another way before we figure out key value pairs
	//var colors2 map[string]string

	//Third way
	//colors3 := make(map[string]string)
	//colors3["white"] = "#ffffff"

	//Deleting keys
	//delete(colors3, "white")
	//fmt.Println(colors3)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

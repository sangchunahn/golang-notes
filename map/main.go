package main

import "fmt"

func main() {
	// map is an object with keys and value pairs
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// same as the bottom
	// colors := make(map[int]string)

	// go's empty value is an empty map
	// var colors map[string]string

	// colors[10] = "#ffffff"

	// delete in the map by selecting the map and the key.
	// delete(colors: 10)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

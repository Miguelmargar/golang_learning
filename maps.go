package main

import "fmt"

func runMaps() {

	// Three different ways to initialise a map

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["red"] = "#FF000"
	// colors["green"] = "#00FF00"

	// delete(colors, "red")

	printMap(colors)

	fmt.Println(colors)
}

func printMap(c map[string]string) {

	for key, value := range c {

		fmt.Println("Hex code for", key, "is", value)
	}
}

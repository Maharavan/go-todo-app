package main

import (
	"fmt"
	"sort"
)

type colorval struct {
	color string
	hex   string
}

func main() {

	color := make(map[string]string)
	color = map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}
	color["white"] = "#ffffff"
	//delete(color, "black")
	sortMap(color)
	fmt.Println(color)
	printMap(color)
}

func printMap(color map[string]string) {
	for col, hex := range color {
		fmt.Println("The Hex of color", col, "is", hex)
	}
}

func sortMap(colors map[string]string) {

	// to sort by key/values require struct of slice
	var structcolor []colorval

	for col, hex := range colors {
		structcolor = append(structcolor, colorval{col, hex})
	}
	sort.Slice(structcolor, func(i, j int) bool {
		return structcolor[i].color < structcolor[j].color
	})

	for _, val := range structcolor {
		colors[val.color] = val.hex
	}
	fmt.Println(colors)
}

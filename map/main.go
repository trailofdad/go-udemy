package main

import "fmt"

func main() {
	// ways to declare maps
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#6ba539",
	}

	colors["white"] = "#ffffff"

	delete(colors, "red")

	fmt.Println(colors)
}

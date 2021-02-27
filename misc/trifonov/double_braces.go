package main

import (
	"fmt"
)

func main() {
	var set map[string]struct{}
	// Initialize the set
	set = make(map[string]struct{})

	// Add some values to the set:
	set["red"] = struct{}{}
	set["blue"] = struct{}{}

	// Check if a value is in the map:
	_, ok := set["red"]
	fmt.Println("Is red in the map?", ok)
	_, ok = set["green"]
	fmt.Println("Is green in the map?", ok)
}

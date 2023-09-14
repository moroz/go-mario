package main

import "fmt"

func getHeight(min, max int) int {
	for {
		var input int
		fmt.Print("Height: ")
		fmt.Scanf("%d", &input)
		if input >= min && input <= max {
			return input
		}
	}
}

func main() {
	height := getHeight(1, 8)

	for y := 1; y <= height; y++ {
		for x := 1; x <= height; x++ {
			if x <= height-y {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

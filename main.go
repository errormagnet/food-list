package main

import (
	"fmt"
	"os"
)

var allRestaurants []restaurants = getFood()

func main() {
	fmt.Println(allRestaurants)
}

func readJson() {
	data, err := os.ReadFile("food.json")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	x := string(data)

	fmt.Println(x)
}

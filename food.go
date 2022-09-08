package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type restaurants struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Position    int        `json:"position"`
	Description string     `json:"description"`
	Images      []string   `json:"images"`
	SubItems    []subItems `json:"sub-items"`
}

type subItems struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Position      int      `json:"position"`
	Price         string   `json:"price"`
	Consumable    string   `json:"consumable"`
	Cuisine_name  string   `json:"cuisine_name"`
	Category_name string   `json:"category_name"`
	Discount      discount `json:"discount"`
	Tags          []string `json:"tags"`
}

type discount struct {
	Type   string `json:"type"`
	Amount string `json:"amount"`
}

func getFood() []restaurants {
	f, rfErr := os.ReadFile("food.json")

	if rfErr != nil {
		fmt.Println("Error: ", rfErr)
		os.Exit(1)
	}

	var payload []restaurants
	umErr := json.Unmarshal(f, &payload)

	if umErr != nil {
		fmt.Println("Error during Unmarshal: ", umErr)
		os.Exit(1)
	}

	return payload

}
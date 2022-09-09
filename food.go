package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type food struct {
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

func getFood(fileName string) []food {
	f, rfErr := os.ReadFile(fileName)

	if rfErr != nil {
		fmt.Println("Error: ", rfErr)
		os.Exit(1)
	}

	var payload []food
	umErr := json.Unmarshal(f, &payload)

	if umErr != nil {
		fmt.Println("Error during Unmarshal: ", umErr)
		os.Exit(1)
	}

	return payload

}

func (f food) print() {
	fmt.Println(f)
}

func getFoodIndex(af []food, fi string) int {
	for i, f := range af {
		if f.Id == fi {
			return i
		}
	}
	return -1
}

func removeFoodAt(af []food, fi int) []food {
	ef := food{}
	af[fi] = af[len(af)-1]
	af[len(af)-1] = ef
	return af[:len(af)-1]
}

func saveToDb(af []food, fileName string) error {
	afj, err := json.Marshal(af)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return os.WriteFile(fileName, afj, 0666)
}

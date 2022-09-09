package main

import (
	"os"
	"testing"
)

func TestSaveToDbAndGetFood(t *testing.T) {
	os.Remove("_deckTesting")

	f := getFood("food.json")
	saveToDb(f, "_deckTesting")

	tf := getFood("_deckTesting")

	if len(tf) < 1 {
		t.Errorf("Expected an array of values but got array of length %v", len(tf))
	}

	os.Remove("_deckTesting")
}

func TestGetFoodIndex(t *testing.T) {
	mf := mf

	mi := "94304"

	var index int

	for i, f := range mf {
		if f.Id == mi {
			index = i
			break
		}
	}
	if index == 0 {
		t.Errorf("Expected index to be truthy")
	}
}

func TestRemoveFoodAt(t *testing.T) {
	mf := mf
	tl := len(mf)
	ef := food{}

	mf[0] = mf[len(mf)-1]
	mf[len(mf)-1] = ef
	mf = mf[:len(mf)-1]

	if len(mf) != tl-1 {
		t.Errorf("Expected new length of food to be one less than original but got %v", len(mf))
	}
}

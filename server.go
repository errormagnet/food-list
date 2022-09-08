package main

import (
	"net/http"

	"github.com/labstack/echo"
)

var allFood []food = getFood()

func registerRoutes() {
	e := echo.New()
	e.GET("/food", getAllFood)
	e.GET("/food/:id", getFoodWithId)
	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	e.DELETE("/food/:id", removeFood)
	e.Logger.Fatal(e.Start(":3000"))
}

func getAllFood(c echo.Context) error {
	return c.JSON(http.StatusOK, allFood)
}

func getFoodWithId(c echo.Context) error {
	pid := c.Param("id")

	for _, food := range allFood {
		if food.Id == pid {
			return c.JSON(http.StatusOK, food)
		}
	}
	return c.String(http.StatusNotFound, string("Food with id "+pid+" does not exist"))
}

func removeFood(c echo.Context) error {
	pid := c.Param("id")
	di := getFoodIndex(allFood, pid)

	if di != -1 {
		allFood = removeFoodAt(allFood, di)
		return c.String(http.StatusOK, "Food with id "+pid+"removed")
	}
	return c.String(http.StatusNotFound, "Food with id "+pid+" does not exist")
}

func createFood(c echo.Context) error {

}

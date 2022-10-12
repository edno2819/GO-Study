package main

import "github.com/labstack/echo/v4"

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func baseCreateCars() {
	cars = append(cars, Car{Name: "ferrari", Price: 100})
	cars = append(cars, Car{Name: "Mercedes", Price: 200})
	cars = append(cars, Car{Name: "Porsh", Price: 300})

}

func main() {
	baseCreateCars()

	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)

	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCars(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	return c.JSON(200, cars)
}

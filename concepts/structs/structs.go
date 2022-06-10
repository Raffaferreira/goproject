package structs

import "fmt"

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"-"`
}

func (c Car) drive() {
	fmt.Println(c.CarName, "Andou")
}

func (c *Car) driveAt() {
	c.CarName = "Fusion"
	fmt.Println(c.CarName, "Andou em direção a..")
}

type CarYearType int

func MainPonteiros() {

	var x CarYearType
	x = 2021
	fmt.Println(x)

	carl1 := Car{
		CarName: "Onix Plus",
		CarYear: 2021,
	}

	carl1.drive()
	fmt.Println("Carro 1 : ", carl1.CarName)

	carl1.driveAt()
	fmt.Println("Carro 2 : ", carl1.CarName)
}

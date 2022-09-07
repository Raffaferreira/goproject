package main

import (
	"example/main/concepts/api"
	mapitem "example/main/concepts/map"
	"fmt"
)

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

func main() {
	mapitem.MainMap()
}

func leftover() {
	x := 10
	fmt.Println(&x)

	api.MainApi()
	// res := abc(&x)
	// fmt.Println(res)

	// array := [10]int{8, 3, 112, 90, 61}
	// fmt.Println(array)

	// slice := make([]int, 5)
	// slice[0] = 1
	// fmt.Println(slice)

	// slice = append(slice, 1, 2, 3, 4, 5)
	// fmt.Println(slice)

	// m := make(map[string]int)

	// m["Fraulyn"] = 24
	// m["Giovana"] = 19
	// fmt.Println(m["Fraulyn"])
}

func whiteloop() int {
	n := 1
	for n < 5 {
		n *= 2
	}

	fmt.Println(n)

	return n
}

func abc(a *int) int {
	return *a
}

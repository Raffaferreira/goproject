package main

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

type Car struct {
	CarName string
	CarYear int
}

func (c Car) drive() {
	fmt.Println(c.CarName, "Andou")
}

type CarYearType int

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	x := 10
	fmt.Println(&x)
	res := abc(&x)
	fmt.Println(res)

	a, b, err := nome("Rafael", 9)
	if err != nil {
		panic("Panico")
	}

	fmt.Println(a, b)

	array := [10]int{8, 3, 112, 90, 61}
	fmt.Println(array)

	slice := make([]int, 5)
	slice[0] = 1
	fmt.Println(slice)

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)

	m := make(map[string]int)

	m["Fraulyn"] = 24
	m["Giovana"] = 19
	fmt.Println(m["Fraulyn"])

	// var x CarYearType
	// x = 2021

	carl := Car{CarName: "Onix Plus", CarYear: 2021}
	fmt.Println(carl)
	carl.drive()

	result := somaTudo(2, 3, 6, 7, 8, 9)
	fmt.Println(result)

	resultado := func(x ...int) func() int {
		valor := 0

		for _, v := range x {
			valor += v
		}

		return func() int {
			return valor * valor
		}
	}

	fmt.Println(resultado(54, 54, 54, 54)())

	array2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	valor := 5
	index := binaria(array2, valor)
	if index == -1 {
		fmt.Errorf("Valor %v não encontrado", valor)
	}

}

func whiteloop() int {
	n := 1
	for n < 5 {
		n *= 2
	}

	fmt.Println(n)

	return n
}

func binaria(array []int, valor int) int {
	// l := 0
	// h := len(array) - 1

	// for l < 5 {
	// 	if array[medio] > n {
	// 		h = m - 1
	// 	}

	// 	if array[medio] < n {
	// 		l = m + 1
	// 	}
	// }

	return 1
}

func abc(a *int) int {
	return *a
}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}

	return resultado
}

func nome(a string, b int) (string, int, error) {
	if b > 10 {
		return "", 0, errors.New("deu ruim")
	}

	return a, b, nil
}

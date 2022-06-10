package arrays

import (
	"fmt"
	"strings"
)

func MainArrays() {
	array2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	valor := 5
	index := binaria(array2, valor)
	if index == -1 {
		fmt.Errorf("Valor %v não encontrado", valor)
	}
}

func binaria(array []int, valor int) int {
	//codigo errado
	l := 0
	h := len(array) - 1

	for l < 5 {
		if array[h] > h {
			h = l - 1
		}

		if array[h] < h {
			l = h + 1
		}
	}

	return 1
}

func SecondMain() {

	bookings := []string{"Rafael", "Karol", "Heloisa"}
	for index, bookings := range bookings {
		var names = strings.Fields(bookings)
		fmt.Println(index)
		fmt.Println(names)
	}

	nomes := []string{"Rafael", "Karol", "Heloisa"}
	firstNames := []string{}
	for _, nome := range nomes {
		var names = strings.Fields(nome)
		firstNames = append(firstNames, names[0])
	}

	fmt.Println(firstNames)

}

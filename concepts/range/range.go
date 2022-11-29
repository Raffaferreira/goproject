package rangepackage

import "fmt"

func Main() {
	var evens = [5]int{2, 4, 6, 8, 10}

	// using index returned from range
	for index := range evens {
		fmt.Printf("evens[%d] = %d \n", index, evens[index])
	}

	var str = "Golang"

	// using index returned from range
	for index := range str {
		fmt.Printf("str[%d] = %d \n", index, str[index])
	}
}

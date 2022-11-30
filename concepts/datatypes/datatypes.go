package datatypes

import "fmt"

func Main() {
	integers()
	floatPointNumbers()
	constdatatype()
}

func constdatatype() {
	const c bool = true
	fmt.Println(c)
}

func integers() {
	fmt.Println(-459)

	var absoluteZero int = -459
	fmt.Println(absoluteZero)

	sum := 116 - 68
	fmt.Println(sum)

}

func floatPointNumbers() {
	fmt.Println(-459.67)
	absoluteZero := 459.67
	fmt.Println(absoluteZero)

	var sum = 564.0 + 365.24
	fmt.Println(sum)
}

package multipleparameters

import "fmt"

func MainMultiple() {

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

}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}

	return resultado
}

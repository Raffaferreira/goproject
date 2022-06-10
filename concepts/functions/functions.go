package functions

import "fmt"

func subtracao(a int, b int) (result int) {
	result = a - b
	return result
}

func divisao(a int, b int) (int, string) {
	return a / b, "Divisao completa!"
}

func somaTudo(valores ...int) int {
	resultado := 0

	for _, v := range valores {
		resultado += v
	}

	return resultado
}

type Pessoa struct {
	Nome string
}

func (p Pessoa) andar() {
	fmt.Println("A pessoa ", p.Nome, " está andando...!")
}

func MainFunctions() {
	resultadoSoma := subtracao(5, 2)
	fmt.Println(resultadoSoma)

	resultadoDivisao, text := divisao(4, 2)
	fmt.Println(resultadoDivisao, text)

	resultadoSomaTudo := somaTudo(3, 5, 92, 39, 55)
	fmt.Println(resultadoSomaTudo)

	resultadoFuncaoAnonima := func(x ...int) func() int {
		resultado := 0

		for _, v := range x {
			resultado += v
		}

		return func() int {
			return resultado * resultado
		}
	}

	fmt.Println(resultadoFuncaoAnonima(3, 5, 92, 39, 55)())

	pessoa := Pessoa{
		Nome: "Rafael Ferreira",
	}

	pessoa.andar()
}

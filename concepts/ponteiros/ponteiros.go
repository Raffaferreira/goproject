package ponteiros

import "fmt"

func MainPonteiros() {
	a := 10

	fmt.Println(&a)

	var ponteiro *int = &a

	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)

	*ponteiro = 50

	fmt.Println(*ponteiro)
	fmt.Println(a)

	variavel := 10
	teste(&variavel)
	fmt.Println(variavel)
}

func teste(a *int) int {
	*a = 200
	return *a
}

func MainPonteirosSecond() {
	variavel := 10
	teste(&variavel)
	fmt.Println(variavel)
}

func MainPonteirosThird() {
	animal := Animal{
		Name: "Chavinho",
	}

	animal.andou()
}

type Animal struct {
	Name string
}

func (a Animal) andou() {
	a.Name = "Wilson"
	fmt.Println(a.Name, "andou...")
}

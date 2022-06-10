package exceptions

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func MainExceptions() {

	//teste 1
	res, err := http.Get("http://google.com.br")

	if err != nil {
		log.Fatal("error : ", err.Error())
	}

	fmt.Println(res.Header)

	//teste 2
	response, erro := soma(5, 6)

	if erro != nil {
		log.Fatalln("Erro na soma dos valores : ", erro.Error())
	}

	fmt.Println(response)

	//teste 3 - blank identifier
	response2, _ := soma(5, 6)
	fmt.Println(response2)
}

func SecondMain() {
	a, b, err := nome("Rafael", 9)
	if err != nil {
		panic("Panico")
	}

	fmt.Println(a, b)
}

func nome(a string, b int) (string, int, error) {
	if b > 10 {
		return "", 0, errors.New("deu ruim")
	}

	return a, b, nil
}

func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Soma maior que 10")
	}

	return res, nil
}

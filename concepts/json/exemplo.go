package jsonpack

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"-"`
}

func JsonExemplo() {
	// ###### Convertendo uma struct para JSON e para um array de bytes, e convertendo um JSON para uma struct
	carl1 := Car{
		CarName: "Onix Plus",
		CarYear: 2021,
	}

	result, _ := json.Marshal(carl1)
	fmt.Println("Marshal converting to slice of bytes : ", result)
	fmt.Println("Marshal converting to string : ", string(result))

	//Convertendo de um array de json, para a estrutura
	j := []byte(`{"car":"BMW", "year": 2022}`)

	var car2 Car

	json.Unmarshal(j, &car2)

	fmt.Println(car2.CarName)

}

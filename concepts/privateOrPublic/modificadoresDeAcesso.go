package modificadoresDeAcesso

import (
	"example/main/exemplo"
	exemplopack "example/main/exemplo"
	"fmt"
)

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"-"`
	Casa    exemplo.Casa
}

func MainExemplo() {
	exemplopack.PrintExemplo()

	carro := Car{
		CarName: "Onix Plus",
		CarYear: 2021,
		Casa: exemplopack.Casa{
			Cor:    "Azul",
			Numero: 98,
			//endereco nao esta acessivel
		},
	}

	//metodo da casa acessivel, pois está com letra inicial maiscula
	carro.Casa.ConstruirCasa()

	casa := exemplo.Casa{
		Cor:    "Branco",
		Numero: 11,
	}

	fmt.Println(casa)
}

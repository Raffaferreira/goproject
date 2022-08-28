package structs

import "fmt"

type Casa struct {
	Cor      string
	Numero   int
	endereco string
}

//Metodo acessivel, pois contem Uppercase da letra inicial
func (cs Casa) ConstruirCasa() {
	fmt.Println("Construindo casa", cs)
}

func PrintExemplo() {
	println("oiiiiiiii")
}

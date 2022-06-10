package inputs

import "fmt"

const teste = 40

var conferenceName = "Go Conference"

func MainInputs() {
	fmt.Printf("Welcome to %v starting programming from scratch \n", conferenceName)
	var bookings = []string{}
	var firstName string

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	bookings = append(bookings, "Olá, sou o primeiro elemento")

	fmt.Printf("Olá %v \n ", firstName)

	fmt.Println(bookings, " Tamanho -> ", len(bookings))
}

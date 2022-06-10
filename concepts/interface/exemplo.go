package interfacepack

type vehicle interface {
	start() string
}

func exemploCarro(car vehicle) {

}

func (c Car) start() string {
	return "Startou"
}

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"-"`
}

func InterfaceExemplo() {
	car := Car{
		CarName: "Tracker Premier",
		CarYear: 2022,
	}

	exemploCarro(car)
}

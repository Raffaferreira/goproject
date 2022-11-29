package general

import (
	"log"
	"reflect"
)

type Teste struct {
	A bool
	B bool
	C bool
}

func main() {
	v := new(Teste)
	metavalue := reflect.ValueOf(v).Elem()

	for _, name := range []string{"A", "B", "C"} {
		field := metavalue.FieldByName(name)
		if field == (reflect.Value{}) {
			log.Printf("Field %s not exist in struct", name)
			//output 2009/11/10 23:00:00 Field Z not exist in struct
		}
	}
}

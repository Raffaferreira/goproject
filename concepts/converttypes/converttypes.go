package converttypes

import (
	"fmt"
	"reflect"
	"strconv"
)

func Main() {
	convertStringToInt()
}

func convertStringToInt() {
	text := "100"
	intvar, err := strconv.Atoi(text)
	fmt.Println(intvar, err, reflect.TypeOf(intvar)) //output 100 <nil> int

	strVar := "100"
	intVar2, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar2, err, reflect.TypeOf(intVar2))

	text2 := "100"
	intvalue := 0
	_, err2 := fmt.Sscan(text2, &intvalue)
	fmt.Println(intvalue, err2, reflect.TypeOf(intvalue))
}

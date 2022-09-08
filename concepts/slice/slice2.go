package slice

import (
	"fmt"
	"reflect"
)

//Declaring Slice using Make specifying the length and the capacity.
func Slice2() {
	var intSlice = make([]int, 10)
	var strSlice = make([]string, 10, 20)

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

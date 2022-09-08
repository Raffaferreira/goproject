package slice

import "fmt"

//Accessing slice by referring the index number
func Slice5() {
	var intSlice = []int{10, 20, 30, 40}
	fmt.Println(intSlice[0])
	fmt.Println(intSlice[1])
	fmt.Println(intSlice[0:4])
}

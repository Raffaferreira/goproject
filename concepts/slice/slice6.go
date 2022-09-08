package slice

import "fmt"

//Updating a specific index of slice
func Slice6() {
	var SliceCountries = []string{"India", "Canada", "Japan"}
	fmt.Println(SliceCountries)

	SliceCountries[2] = "Germany"
	fmt.Println(SliceCountries)
}

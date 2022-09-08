package slice

import "fmt"

func Slice7() {
	var SliceCountries = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(SliceCountries)

	SliceCountries = RemoveIndex(SliceCountries, 3)
	fmt.Println(SliceCountries)
}

func RemoveIndex(source []string, index int) []string {
	return append(source[:index], source[index+1:]...)
}

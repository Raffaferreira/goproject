package general

import (
	"fmt"
	"strings"
)

func Main() {
	str1 := " Hello,   World  ! "
	fmt.Println(standardizeSpaces(str1))
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

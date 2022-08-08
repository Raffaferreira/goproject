package shortvariabledeclaration

import "fmt"

func Main() {

	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)

	k = 4
	fmt.Println(k)

	k = 4
	fmt.Println(&k) //memory address
}

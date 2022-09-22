package operators

import (
	"fmt"
	"math"
)

//	.	x + y	Sum of x and y
//	.	x - y	Difference of x and y
//	.	-x		Changed sign of x
//	.	+x		Identity of x
//	.	x * y	Product of x and y
//	.	x / y	Quotient of x and y
//	.	x % y	Remainder of x / y
func Operators() {
	integers()
	floats()
	subtraction()
	unusedDataTypes()
	unaryOperation()
	multiplication()
	division()
}

func integers() {
	a := 88
	b := 103
	fmt.Println((a + b))

	c := -36
	d := 25
	fmt.Println(c + d)
}

func floats() {
	e := 5.5
	f := 2.5
	fmt.Println((e + f))    // output 8
	fmt.Printf("%.2f", e+f) // It will format to two decimal places : output 8.00
}

func subtraction() {
	g := 75.67
	h := 32.0

	fmt.Println(g - h)
}

// In go, we can only use operators on the same data type. We can't add an "int" and a "float64"
func unusedDataTypes() {
	i := 7
	j := 7.0
	fmt.Println(i)
	fmt.Println(j)
	// Using operators on different data types will result in a compiler error
	//fmt.Println(i + j) // i + j (mismatched types int and float64) //
}

func unaryOperation() {
	i := 3.3
	fmt.Println(+i) //output 3.3

	j := -19
	fmt.Println(-j)

	k := 3.3
	fmt.Println(-k)
}

func multiplication() {
	k := 100.2
	j := 10.2
	result := k * j
	fmt.Println(result)
}

func division() {
	m := 80
	n := 6
	fmt.Println(m / n) // output 13

	s := 80
	t := 6
	r := float64(s) / float64(t)
	fmt.Println(r) // Output 13.333333333333334
}

//The % operator is the modulo, which returns the remainder rather than the quotient after division.
func modulo() {
	o := 85
	p := 15

	fmt.Println(o % p)

	q := 36.0
	r := 8.0
	s := math.Mod(q, r)
	fmt.Println(s)
}

package slice

import "fmt"

func Slice4() {
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A :", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	a = append(a, 30)
	fmt.Print("Slice A after appending data :", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
}

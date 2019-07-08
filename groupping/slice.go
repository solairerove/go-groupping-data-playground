package groupping

import (
	"fmt"
)

// SliceExample tbd
func SliceExample() {

	// x := type{values} // composite literal

	x := []int{1, 2, 3, 5}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(len(x))

	fmt.Println("for loop")
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println("for loop _")
	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Println("slicing")
	fmt.Println(x[:])
	fmt.Println(x[1:])

	fmt.Println("append")
	fmt.Println(append(x, 42))

	y := []int{30, 50}
	fmt.Println("append", y)
	fmt.Println(append(x, y...))
}

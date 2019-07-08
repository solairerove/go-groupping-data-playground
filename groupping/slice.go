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
	x = append(x, y...)
	fmt.Println(x)

	// efficient compilation
	// efficient execution
	// "ease" of programming
	fmt.Println("delete due append")
	fmt.Println(append(x[:1], x[2:]...))

	// increase capacity
	fmt.Println("make")
	x = make([]int, 10, 15)
	fmt.Println(x)

	x[9] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

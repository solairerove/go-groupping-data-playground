package groupping

import (
	"fmt"
)

// ArraysExample tbd
func ArraysExample() {

	var x [5]int
	fmt.Println(x)

	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}

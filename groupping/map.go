package groupping

import (
	"fmt"
)

// MapExample tbd
func MapExample() {

	x := map[string]int{
		"james":  32,
		"mikita": 23,
	}
	fmt.Println(x)
	fmt.Println(x["mikita"])
	// why zero value?
	fmt.Println(x["mikita2"])

	// comma ok idiom
	v, ok := x["mikita2"]
	fmt.Println(v, ok)

	if y, mkay := x["mikita"]; mkay {
		fmt.Printf("printed only if map has that key %v. here's value %v", "[mikita]", y)
	}
}

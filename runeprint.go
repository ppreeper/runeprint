package main

import (
	"fmt"
)

//PowInt return intpower
func PowInt(b int, p int) int64 {
	v := int64(1)
	for i := 1; i <= p; i++ {
		v *= int64(b)
	}
	return v
}

func main() {
	var correct = "\u2713"
	var incorrect = "\u2715"
	fmt.Printf("%s\t%s\n", correct, incorrect)
	for i := 0; i < int(PowInt(2, 20)); i++ {
		rs := rune(i)
		fmt.Printf("int: %d\trune: %U\t%s\n", i, rs, string(rs))
	}
	// for i := 128513; i < 128591; i++ {
	// 	rs := rune(i)
	// 	fmt.Printf("int: %d\trune: %U\t%s\n", i, rs, string(rs))
	// }
}

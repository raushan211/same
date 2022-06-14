package main

import (
	"fmt"
	"strings"
)

func main() {
	var word1 = []string{"ab", "c"}
	var word2 = []string{"a", "bc"}
	// var a string
	// var b string

	// for i := 0; i < len(word1); i++ {
	// 	a = a + word1[i]

	// }
	// for j := 0; j < len(word2); j++ {
	// 	b = b + word2[j]

	// }

	// if a == b {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }

	fmt.Println(strings.Join(word1, "") == strings.Join(word2, ""))
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(hamming("Hello, 世界", "Hello, e界"))
}

// make the string comparison
func hamming(str1 string, str2 string) int {
	
	runesi1 := []rune(str1)
	runesi2 := []rune(str2)
	if len(runesi1) != len(runesi2) {
		return -1
	}
	result := 0
	for position, vRune := range runesi1 {{
		if vRune != runesi2[position] {
			result++
		}
	}

	return result
}

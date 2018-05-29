package main

import (
	"fmt"
	"strings"
)

/* PositiveSum */
func PositiveSum(numbers []int) int {

	if numbers == nil || len(numbers) < 0 {
		return 0
	}
	var sum = 0
	for _, i := range numbers {
		if i > 0 {
			sum += i
		}
	}
	return sum
}

func bandNameGenerator1(word string) string {
	first := word[:1]
	last := word[len(word)-1:]

	if first != last {
		return "The " + strings.Title(word)
	}

	return strings.Title(word) + word[1:]
}

/**
 * 复杂版本
 */
func BandNameGenerator(word string) string {
	// chars := []rune(word)
	finalStr := ""
	if word[0] == word[len(word)-1] {
		finalStr := splitAndUpper(word)
		return finalStr
	} else {
		finalStr := titleAndThe(word)
		return finalStr
	}
	// strs
	// strings.to
	// if {

	// }

	return finalStr
}

func splitAndUpper(word string) string {
	finalStr := word
	r := []rune(finalStr)
	for i := 0; i < len(r); i++ {
		fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
		finalStr += string(r[i])
	}
	// for i := 1; i < len(strs); i++ {
	// 	finalStr += strs[i]
	// }
	return strings.Title(finalStr)
}

func titleAndThe(word string) string {
	finalStr := "the "
	finalStr += word
	return strings.Title(finalStr)
}

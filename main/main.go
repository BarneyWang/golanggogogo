package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
	// PositiveSum(nil)
	// PositiveSum([]int{1, 2, 3})
	// fmt.Println(BandNameGenerator("wangd"))
	// PositiveSum([]int{2, 33, 44})
	// PositiveSum(nil)o
	str := "1213"
	fmt.Println(str[:3])
	// chars := []rune(str)
	// fmt.Println(chars[0] == chars[len(chars)-1])
	fmt.Println(BandNameGenerator("alaska"))
	fmt.Println(RepeatStr(5, "hello"))
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	max := 0
	m := strings.Split(a, "")
	for j := 0; j < len(m); j++ {
		k, err := strconv.Atoi(m[j])
		if err != nil {
			panic("Строка чисел содержеит букву")
		}
		if k > max {
			max = k
		}
	}
	fmt.Println(max)
}

package main

import "fmt"


func supperAdd(numbers ...int) int {
	// range : array 에 loop 를 해줌
	total := 0
	// forEach 와 비슷함
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := supperAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(result)
}

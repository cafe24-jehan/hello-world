package main

import "fmt"

func main() {
	// array의 제한 없이 쓰고싶을때 쓸 수 있는 data type = slice
	names := [5]string{"jieun", "insik", "down"}
	fmt.Println(names)
}

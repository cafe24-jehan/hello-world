package main

import "fmt"

func main() {
	// array의 제한 없이 쓰고싶을때 쓸 수 있는 data type = slice
	names := []string{"jieun", "insik", "down"}
	names = append(names, "lala") // 하나의 slice와 값을 값아서 새로운 array를 return 한다.
	
	fmt.Println(names)
}

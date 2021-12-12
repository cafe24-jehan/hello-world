package main

import "fmt"

func main() {
	a := 2
	b := &a // 메모리 주소를 저장
	*b = 20 // b는 a를 바라보는 포인터이기때문에 *b를 이용해 a의 값을 바꿈
	fmt.Println(a)
}

// 메모리 주소를 보는 법 = &
// 메모리에 저장된 값을 보는 법 = *
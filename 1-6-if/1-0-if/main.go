package main

import "fmt"


func canIDrink(age int) bool {
	// if 조건을 체크하기전에 invariable 을 생성할 수 있다
	// 해당 if else에서만 사용할 수 있음
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} 
	return true
}

func main() {
	fmt.Println(canIDrink(15))
}

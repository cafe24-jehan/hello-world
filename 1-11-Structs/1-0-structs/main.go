package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}


func main() {
	
	favFood := []string{"kimchi", "ramen"}
	// 이렇게 작성할 경우 field가 무엇인지 정확히 인지하기 어려울 수 있음
	// nico := person{"nico", 18, favFood}

	// struct 가 없거나 보이지 않는 경우에도 명확히 알 수 있음
	nico := person{name:"nico", age:18, favFood: favFood}
	fmt.Println(nico)
}

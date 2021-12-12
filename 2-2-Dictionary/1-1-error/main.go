package main

import (
	"fmt"
	"naked/accounts"
)


func main() {
	account := accounts.NewAccount("jieun")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err) Println 을 호출하고 프로그램을 종료시킴
	}
	fmt.Println(account.Balance())

}

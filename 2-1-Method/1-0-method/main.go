package main

import (
	"fmt"
	"naked/accounts"
)


func main() {
	account := accounts.NewAccount("jieun")
	account.Deposit(10)
	fmt.Println(account.Balance())
}

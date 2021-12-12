package accounts

import "errors"

// Account struct
type Account struct {
	owner string
	balance int
}

var errNoMoney =  errors.New("Can't withdraw")

// NewAccount
// function을 만들어 object를 return 시킨다
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account // 실제 메모리 address 를 리턴함. 복사본을 리턴하는 것이 아님
}

// 기본적으로 method는 리시버를 가지고 있음
// 아래 func 에서 리시버는 a 이다. a의 타입은 Account
// 리시버 작성 시 유의해야 할 사항 : 구조체의 첫글자를 따서 소문자로 작성
// Deposit x amount your account
// pointer receiver - account나 receiver를 복사하지 말고 실제 receiver를 달라고 하는 것
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int{
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	// a.balance -= amount
	if a.balance < amount {
		return errNoMoney
	} 
	a.balance -= amount
	return nil // error value : error/nil
}
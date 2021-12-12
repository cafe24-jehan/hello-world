package accounts

// Account struct
type Account struct {
	owner string
	balance int
}

// NewAccount
// function을 만들어 object를 return 시킨다
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account // 실제 메모리 address 를 리턴함. 복사본을 리턴하는 것이 아님
}
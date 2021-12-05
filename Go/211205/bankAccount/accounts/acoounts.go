package accounts

import (
	"errors"
	"fmt"
)

//go는 constructor가 없기 때문에 function으로 construct 하거나 struct를 만들어서 씀
//Account strunct
type Account struct {
	// 대문자로 해야  export됨
	/*
		소문자 시작 : private 로 인식
		대문자 시작 : public 으로 인식
	*/
	owner   string
	balance int
}

//같은 에러 전역적 사용시
var errNoMoney = errors.New("계좌에 인출할 돈이 부족합니다.")

//new account create
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

/*
	balance 증가를 위해 function 말고 method를 사용한다
	method는 fucn 바로 뒤에 receiver&type 를 넣는 것만으로 메소드 구현 가능
	receiver는 a(마음대로 지정가능) & a의 type은 Acoount인 상태이다
	##receiver의 규칙 - struct의 첫 글자 따서 소문자로 지어야 함(Account의 a)

*/
// Deposit X amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw X amount from yout account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney //errors.New("계좌에 인출할 돈이 부족합니다.")
	}
	a.balance -= amount
	return nil //error사용시 여기도 리턴해줘야 됨 nil = null
}

//ChangeOwner of ther accouont
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account\nHas : ", a.Balance())
}

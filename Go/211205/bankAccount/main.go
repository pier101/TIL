package main

import (
	"fmt"
	_ "log"

	"github.com/pier101/leango-nomade/bankAccount/accounts"
)

func main() {
	//출력은 되지만 누구나 접근해 수정할 수 있는 문제가 있음
	// account := banking.Account{Owner: "dongwook", Balance: 1000}
	// fmt.Println(account)

	account := accounts.NewAccount("wook")
	account.Deposit(200)
	fmt.Println(account)
	err := account.Withdraw(300)
	if err != nil {
		//log.Fatalln(err)
		//Fatalln 프로그램을 종료기켜줌
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	account.ChangeOwner("dong")
	fmt.Println(account.Owner())
	//
}

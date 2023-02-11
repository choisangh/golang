package banking

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int // 스트럭트 속성도 첫글자 대문자여야 public
}

var errNoMoney = errors.New("Can't withdraw") //에러 변수

// NewAccount creates Account = construtor
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account //생성 후 바로 리턴 포인터
}

func (a Account) Balance() int {
	return a.balance
}

func (a Account) Owner() string {
	return a.owner
}

func (a *Account) Deposit(amount int) {
	//point reciever
	//Method - Go는 복사본을 만들기 때문 *Account로 해야 적용
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if (a.balance - amount) < 0 {
		return errNoMoney
	}
	a.balance -= amount
	return nil //nil=none

}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \n Has:", a.Balance())
}

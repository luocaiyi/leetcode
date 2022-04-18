package _241_design_an_atm_machine

import (
	"fmt"
	"testing"
)

func TestATM(t *testing.T) {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	// 0 0 1 0 1
	fmt.Printf("%+v\n", atm.Withdraw(600))
	atm.Deposit([]int{0, 1, 0, 1, 1})
	// -1
	fmt.Printf("%+v\n", atm.Withdraw(600))
	// 0 1 0 0 1
	fmt.Printf("%+v\n", atm.Withdraw(550))
}

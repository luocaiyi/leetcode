package _241_design_an_atm_machine

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

type ATM struct {
	cashNum  []int
	cashType []int
}

func Constructor() ATM {
	cashNum := []int{0, 0, 0, 0, 0}
	cashType := []int{20, 50, 100, 200, 500}
	return ATM{cashNum, cashType}
}

func (a *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		a.cashNum[i] += banknotesCount[i]
	}
}

func (a *ATM) Withdraw(amount int) []int {
	ans := make([]int, 5)
	for i := 4; i >= 0 && amount > 0; i-- {
		if a.cashNum[i] == 0 || a.cashType[i] > amount {
			continue
		}
		ans[i] = min(a.cashNum[i], amount/a.cashType[i])
		amount -= ans[i] * a.cashType[i]
	}
	if amount != 0 {
		return []int{-1}
	}
	for i := range ans {
		a.cashNum[i] -= ans[i]
	}
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

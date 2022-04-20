package _281_subtract_the_product_and_sum_of_digits_of_an_integer

func subtractProductAndSum(n int) int {
	a := 1
	b := 0
	for n != 0 {
		x := n % 10
		a *= x
		b += x
		n /= 10
	}
	return a - b
}

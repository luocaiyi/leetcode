package _0_powx_n

// 非递归 - 位运算
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow float64 = 1
	for n != 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}

// 递归
//func myPow(x float64, n int) float64 {
//	if n == 0 { return 1 }
//	if n < 0 {
//		return 1 / myPow(x, -n)
//	}
//	if n % 2 == 1 {
//		return x * myPow(x, n - 1)
//	}
//	return myPow(x * x, n / 2)
//}

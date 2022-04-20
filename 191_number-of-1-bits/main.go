package _91_number_of_1_bits

func hammingWeight(num uint32) int {
	n := 0
	for ; num != 0; n++ {
		num &= num - 1
	}
	return n
}

//func hammingWeight(num uint32) (ones int) {
//	for i := 0; i < 32; i++ {
//		if 1<<i&num > 0 {
//			ones++
//		}
//	}
//	return
//}

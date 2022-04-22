package _588_sum_of_all_odd_length_subarrays

func sumOddLengthSubarrays(arr []int) (sum int) {
	n := len(arr)
	for i, v := range arr {
		leftCount, rightCount := i, n-i-1
		leftOdd := (leftCount + 1) / 2
		rightOdd := (rightCount + 1) / 2
		leftEven := leftCount/2 + 1
		rightEven := rightCount/2 + 1
		sum += v * (leftOdd*rightOdd + leftEven*rightEven)
	}
	return
}

//func sumOddLengthSubarrays(arr []int) (sum int) {
//	n := len(arr)
//	prefixSum := make([]int, n+1)
//	for i, v := range arr {
//		prefixSum[i+1] = prefixSum[i] + v
//	}
//	for start := range arr {
//		for length := 1; start + length <= n; length += 2 {
//			end := start + length - 1
//			sum += prefixSum[end+1] - prefixSum[start]
//		}
//	}
//	return
//}

//func sumOddLengthSubarrays(arr []int) (sum int) {
//	n := len(arr)
//	for start := range arr {
//		for length := 1; start + length <= n; length += 2 {
//			for _, v := range arr[start: start + length] {
//				sum += v
//			}
//		}
//	}
//	return
//}

package __two_sum

func twoSum(nums []int, target int) []int {
	numsSet := map[int]int{}
	for i, x := range nums {
		if j, ok := numsSet[target-x]; ok {
			return []int{j, i}
		}
		numsSet[x] = i
	}
	return nil
}

package _8_4sum

import "sort"

func fourSum(nums []int, target int) (res [][]int) {
	if len(nums) < 4 {
		return
	}
	n := len(nums)
	sort.Ints(nums)
	for first := 0; first < n-3 && nums[first]+nums[first+1]+nums[first+2]+nums[first+3] <= target; first++ {
		if first > 0 && nums[first] == nums[first-1] || nums[first]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for second := first + 1; second < n-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] || nums[first]+nums[second]+nums[n-2]+nums[n-1] < target {
				continue
			}
			// 单指针
			for third, four := second+1, n-1; third < n-1; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				for third < four && nums[third]+nums[four] > target-nums[first]-nums[second] {
					four--
				}
				if third == four {
					break
				}
				if nums[third]+nums[four] == target-nums[first]-nums[second] {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
			// 双指针
			//for third, four := second+1, n-1; third < four; {
			//	if sum := nums[first] + nums[second] + nums[third] + nums[four]; sum == target {
			//		res = append(res, []int{nums[first], nums[second], nums[third], nums[four]})
			//		for third++; third < four && nums[third] == nums[third-1]; third++ {}
			//		for four--; third < four && nums[four] == nums[four+1]; four-- {}
			//	} else if sum < target {
			//		third++
			//	} else {
			//		four--
			//	}
			//}
		}
	}
	return
}

package _5_3sum

import (
	"sort"
)

// 左右指针
func threeSum(nums []int) (res [][]int) {
	if len(nums) < 3 {
		return
	}
	sort.Ints(nums)
	for i := range nums[:len(nums)-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l++
			} else if s > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return
}

// 使用map寻找剩下两位元素
//func threeSum(nums []int) (res [][]int) {
//	if len(nums) < 3 {
//		return
//	}
//	sort.Ints(nums)
//
//	m := make(map[string]struct{})  // 用于结果集去重，结果集用set去重会更好
//	for i, v := range nums[:len(nums)-2] {
//		if i > 0 && v == nums[i-1] {  // 去除重复的v
//			continue
//		}
//		d := map[int]int{}
//		for _, x := range nums[i+1:] {
//			if _, ok := d[x]; !ok {  // 找不到x就将(0-v-x)的值放入
//				d[-v-x] = 1
//			} else {
//				// 统计结果集+去重
//				if _, ok := m[fmt.Sprintf("%d%d%d", v, -v-x, x)]; !ok {
//					m[fmt.Sprintf("%d%d%d", v, -v-x, x)] = struct{}{}
//					res = append(res, []int{v, -v-x, x})
//				}
//			}
//		}
//	}
//	return
//}

//func threeSum(nums []int) (res [][]int) {
//	n := len(nums)
//	sort.Ints(nums)
//
//	// 枚举a
//	for first := 0; first < n; first++ {
//		if first > 0 && nums[first] == nums[first-1] { continue }
//		third := n-1
//		target := -1*nums[first]
//		for second := first+1; second < n; second++ {
//			if second > first+1 && nums[second] == nums[second-1] { continue }
//			for second < third && nums[second] + nums[third] > target { third-- }
//			if second == third { break }
//			if nums[second] + nums[third] == target {
//				res = append(res, []int{nums[first], nums[second], nums[third]})
//			}
//		}
//	}
//	return
//}

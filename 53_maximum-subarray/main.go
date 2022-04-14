package _3_maximum_subarray

// 动态规划
//func maxSubArray(nums []int) int {
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] + nums[i-1] > nums[i] {
//			nums[i] += nums[i-1]
//		}
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	return max
//}

// 分治
func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

// pushUp 这个分治方法类似于「线段树求解最长公共上升子序列问题」的 pushUp 操作
/**
 *  - 首先最好维护的是 iSum，区间 [l,r] 的 iSum 就等于「左子区间」的 iSum 加上「右子区间」的 iSum。
 *  - 对于 [l,r] 的 lSum，存在两种可能，二者取大
 *  	- 要么等于「左子区间」的 lSum
 *  	- 要么等于「左子区间」的 iSum 加上「右子区间」的 lSum
 *  - 对于 [l,r] 的 rSum，同理，二者取大
 *  	- 要么等于「右子区间」的 rSum
 *  	- 要么等于「右子区间」的 iSum 加上「左子区间」的 rSum
 *  - 当计算好上面的三个量之后，就很好计算 [l,r] 的 mSum 了。我们可以考虑 [l,r] 的 mSum 对应的区间是否跨越 m
 *  	- 它可能不跨越 m，也就是说 [l,r] 的 mSum 可能是「左子区间」的 mSum 和 「右子区间」的 mSum 中的一个；
 *  	- 它也可能跨越 m，可能是「左子区间」的 rSum 和 「右子区间」的 lSum 求和。三者取大。
 */
func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSun := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSun, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Status
/**
 *  - lSum 表示 [l,r] 内以 l 为左端点的最大子段和
 *  - rSum 表示 [l,r] 内以 r 为右端点的最大子段和
 *  - mSum 表示 [l,r] 内的最大子段和
 *  - iSum 表示 [l,r] 的区间和
 */
type Status struct {
	lSum, rSum, mSum, iSum int
}

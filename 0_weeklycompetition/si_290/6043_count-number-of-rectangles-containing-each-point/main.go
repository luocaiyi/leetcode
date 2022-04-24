package _043_count_number_of_rectangles_containing_each_point

import "sort"

func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][1] > rectangles[j][1]
	})
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] > points[j][1]
	})
	ans := make([]int, len(points))
	i, n, xs := 0, len(rectangles), []int{}
	for _, p := range points {
		start := i
		for ; i < n && p[1] <= rectangles[i][1]; i++ {
			xs = append(xs, rectangles[i][0])
		}
		if start < i {
			sort.Ints(xs)
		}
		ans[p[2]] = i - sort.SearchInts(xs, p[0])
	}
	return ans
}

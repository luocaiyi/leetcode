package _242_maximum_score_of_a_node_sequence

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)
	g := make([][3]int, n)
	for i := range g {
		g[i] = [3]int{-1, -1, -1}
	}
	var edit func(u, v int)
	edit = func(u, v int) {
		if g[u][0] == -1 || scores[v] > scores[g[u][0]] {
			g[u][0], g[u][1], g[u][2] = v, g[u][0], g[u][1]
		} else if g[u][1] == -1 || scores[v] > scores[g[u][1]] {
			g[u][1], g[u][2] = v, g[u][1]
		} else if g[u][2] == -1 || scores[v] > scores[g[u][2]] {
			g[u][2] = v
		}
	}
	for _, e := range edges {
		edit(e[0], e[1])
		edit(e[1], e[0])
	}
	res := -1
	for _, e := range edges {
		u := e[0]
		v := e[1]
		w := scores[u] + scores[v]
		for i := 0; i < 3; i++ {
			l := g[u][i]
			if l == -1 {
				break
			}
			if l == v {
				continue
			}
			for j := 0; j < 3; j++ {
				r := g[v][j]
				if r == -1 {
					break
				}
				if r == u || r == l {
					continue
				}
				if w+scores[l]+scores[r] > res {
					res = w + scores[l] + scores[r]
				}
			}
		}
	}
	return res
}

//type nb struct {
//	to, s int
//}
//
//func maximumScore(scores []int, edges [][]int) int {
//	g := make([][]nb, len(scores))
//
//	for _, e := range edges {
//		x, y := e[0], e[1]
//		g[x] = append(g[x], nb{y, scores[y]})
//		g[y] = append(g[y], nb{x, scores[x]})
//	}
//
//	for i, vs := range g {
//		sort.Slice(vs, func(i, j int) bool {
//			return vs[i].s > vs[j].s
//		})
//		if len(vs) > 3 {
//			vs = vs[:3]
//		}
//		g[i] = vs
//	}
//
//	ans := -1
//	for _, e := range edges {
//		x, y := e[0], e[1]
//		for _, p := range g[x] {
//			for _, q := range g[y] {
//				if p.to != y && q.to != x && p.to != q.to {
//					ans = max(ans, p.s + scores[x] + scores[y] + q.s)
//				}
//			}
//		}
//	}
//	return ans
//}
//
//func max(a, b int) int {
//	if b > a {return b}
//	return a
//}

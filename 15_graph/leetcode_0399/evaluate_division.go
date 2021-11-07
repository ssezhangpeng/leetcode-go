package leetcode0399

var graph = make(map[string]map[string]float64)

// TODO: 未完成
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 构造有向无环图
	for i, e := range equations {
		graph[e[0]] = map[string]float64{
			e[1]: float64(values[i]),
		}

		graph[e[1]] = map[string]float64{
			e[0]: float64(1.0 / values[i]),
		}
	}

	// 对每一个查询进行 dfs 求值
	var ans []float64
	for _, query := range queries {
		visited := make(map[string]struct{})
		cost := dfs(query[0], query[1], visited)
		ans = append(ans, cost)
	}

	return ans
}

func dfs(src, dst string, visited map[string]struct{}) float64 {
	// 递归出口
	if _, ok := graph[src][dst]; ok {
		return graph[src][dst]
	}

	for neighbor, cost := range graph[src] {
		if _, ok := visited[neighbor]; ok {
			continue
		}

		visited[neighbor] = struct{}{}
		t := dfs(neighbor, dst, visited)
		if t > float64(0) {
			return t * cost
		}
	}
	return float64(-1.0)
}

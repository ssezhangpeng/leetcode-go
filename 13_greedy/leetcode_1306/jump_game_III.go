package leetcode_1306

func canReach(arr []int, start int) bool {
	if arr[start] == 0 {
		return true
	}

	used := make([]bool, len(arr))
	var queue []int

	queue = append(queue, start)
	used[start] = true

	for len(queue) > 0 {
		idx := queue[0]
		queue = queue[1:]
		if idx+arr[idx] < len(arr) && !used[idx+arr[idx]] {
			if arr[idx+arr[idx]] == 0 {
				return true
			}
			queue = append(queue, idx+arr[idx])
			used[idx+arr[idx]] = true
		}

		if idx-arr[idx] >= 0 && !used[idx-arr[idx]] {
			if arr[idx-arr[idx]] == 0 {
				return true
			}
			queue = append(queue, idx-arr[idx])
			used[idx-arr[idx]] = true
		}
	}
	return false
}

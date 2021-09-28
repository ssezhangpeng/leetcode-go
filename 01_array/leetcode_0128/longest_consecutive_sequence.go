package leetcode_0128

func longestConsecutive(nums []int) int {
	exist := make(map[int]bool)
	used := make(map[int]bool)

	for i:=0; i<len(nums); i++ {
		exist[nums[i]] = true
	}

	maxLength := 0
	for i:=0; i<len(nums); i++ {
		if used[nums[i]] {
			continue
		}

		used[nums[i]] = true
		currLength := 1

		for j:=nums[i]+1; exist[j] && !used[j]; j++ {
			used[j] = true
			currLength += 1
		}

		for j:=nums[i]-1; exist[j] && !used[j]; j-- {
			used[j] = true
			currLength += 1
		}

		maxLength = max(maxLength, currLength)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
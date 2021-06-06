package leetcode_0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		m[nums[i]] = i
	}

	for i:=0; i<len(nums); i++ {
		idx, ok := m[target-nums[i]]
		if ok && idx != i {
			return []int{i, idx}
		}
	}
	return nil
}

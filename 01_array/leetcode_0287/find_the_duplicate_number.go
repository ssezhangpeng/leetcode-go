package leetcode_0287

func findDuplicate(nums []int) int {
	// 构造链表，则形成的链表一定是有环的，找到环状入口，即为重复元素
	slow := nums[0]
	fast := nums[nums[0]]

	// slow 每次走一步，fast 每次走两步
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// slow 从开始处，fast 从相遇处，每次都走一步
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

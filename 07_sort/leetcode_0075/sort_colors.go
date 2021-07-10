package leetcode_0075

func sortColors(nums []int) {
	red, blue := 0, len(nums)-1

	for i := 0; i < blue+1; {
		if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			i++
			red++
		} else if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
		} else {
			i++
		}
	}
	return
}

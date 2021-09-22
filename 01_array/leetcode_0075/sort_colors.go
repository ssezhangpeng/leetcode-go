package leetcode_0075

func sortColors(nums []int)  {
	red, blue := 0, len(nums) - 1

	for i:=0; i<=blue; i++ {
		if nums[i] == 0 {
			nums[red], nums[i] = nums[i], nums[red]
			red++
		} else if nums[i] == 2 {
			nums[blue], nums[i] = nums[i], nums[blue]
			blue--
			i--
		}
	}

	return
}

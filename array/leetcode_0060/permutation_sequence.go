package leetcode_0060

import "strconv"

func getPermutation(n int, k int) string {
	arr := make([]int, n)
	for i:=1; i<=n; i++ {
		arr[i-1] = i
	}
	for i:=1; i<k; i++ {
		nextPermutation(arr)
	}

	ans := ""
	for i:=0; i<n; i++ {
		ans += strconv.Itoa(arr[i])
	}
	return ans
}

func nextPermutation(nums []int) {
	len := len(nums)
	if len < 2 {
		return
	}

	// step1: find arr[i]<arr[i+1], from right --> left
	i, j := 0, 0
	for i=len-2; i>=0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i < 0 {
		reverse(nums, 0, len-1)
	}

	//step2: find arr[j] > arr[i], from right --> left
	for j=len-1; j>i; j-- {
		if nums[j] > nums[i] {
			break
		}
	}
	//step3: swap(arr[i], arr[j])
	nums[i], nums[j] = nums[j], nums[i]
	//step4: reverse(arr[i+1], arr[len-1])
	reverse(nums, i+1, len-1)
}

func reverse(nums []int, start, end int) {
	for i,j := start,end; i<j; i,j = i+1,j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

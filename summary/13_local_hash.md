## 原地哈希专题
### 题目汇总
    - leetcode_0041
    - leetcode_0442
    - leetcode_0448


### 使用场景
    - 求第一个缺失的数字
    - 求重复出现的数字


### 知识点
    - 本地Hash

### 重要步骤
    - 1. 将 nums[i] 交换到下标为 nums[i]-1 的位置上
    - 2. 从头开始遍历，遇到 nums[i] 不等于 i + 1 的时候进行处理(nums[i]代表重复的元素; i+1 代表缺失的数字)


### 解题模版
```go
// leetcode_0041
func firstMissingPositive(nums []int) int {
    // 原地 Hash
    size := len(nums)

    for i:=0; i<size; i++ {
        for nums[i] > 0 && nums[i] <= size && nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }

    for i:=0; i<size; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }

    return size + 1
}
```

### 专题总结
## 快速选择专题
### 题目汇总
    - leecode_0215


### 使用场景
    - 求第K大(小)的元素
    - 当然同样的场景下也可以使用小根堆(大根堆)去解决问题


### 知识点
    - 快速选择
    - 分割过程

### 重要步骤
    - 1. 类似快速排序, 递归从左边部分进行快速选取 ｜ 从右边部分进行快速选择
    - 2. 分割函数和快速排序的一模一样


### 解题模版
```go
// leetcode_0215
func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
    q := partition(a, l, r)
    if q == index {
        return a[q]
    } else if q < index {
        return quickSelect(a, q + 1, r, index)
    }
    return quickSelect(a, l, q - 1, index)
}

func partition(nums []int, low int, high int) int {
    // 一趟分割过程
    pivot := nums[low]

    for low < high {
        for low<high && nums[high] >= pivot {
            high--
        }
        nums[low] = nums[high]

        for low<high && nums[low] <= pivot {
            low++
        }
        nums[high] = nums[low]
    }

    nums[low] = pivot

    return low
}
```

### 专题总结
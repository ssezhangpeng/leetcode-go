## 二分图专题
### 题目汇总
    - leetcode_0886


### 使用场景
    - 分组

### 知识点
    - 二分图思想

### 重要步骤
    - 1. 构建图
    - 2. 对图中每个节点作为起点进行DFS染色


### 解题模版
```go
// leetcode_0886
const N = 2010
const M = 2 * 10010
// he数组: 存储是某个节点所对应的边的集合（链表）的头结点
var he = [N]int{}
// e数组: 由于访问某一条边指向的节点；
var e = [M]int{}
// ne数组: 由于是以链表的形式进行存边，该数组就是用于找到下一条边；
var ne = [M]int{}
// color数组: 用于表示节点的颜色
var color = [N]int{}
var idx int

func add(a, b int) {
    e[idx] = b
    ne[idx]=he[a]
    he[a] = idx
    idx++
}

func dfs(u int, colour int) bool {
    // 将节点u染成colour的颜色
    color[u] = colour
    // 遍历所有与节点u相邻的节点，1.如果和u颜色相同，则失败；2.如果还没染色，则尝试染成不同颜色；3. 如果已经染色并且和节点u颜色不同，则符合要求，直接跳过
    for i:=he[u]; i!=-1; i=ne[i] {
        // j为与节点u相邻的节点
        j := e[i]
        if color[j] == color[u] {
            return false
        }
        /*
            colour=0: 未进行染色
            colour=1: 染成红色
            colour=2: 染成黄色
            注意: 3-color为互斥的另外一种颜色
        */
        if color[j] == 0 && !dfs(j, 3-colour) {
            return false
        }
    }
    return true
}

func possibleBipartition(n int, dislikes [][]int) bool {
    // 初始化邻接矩阵
    for i:=0; i<len(he); i++ {
        he[i] = -1
    }
    for _, ds := range dislikes {
        a, b := ds[0], ds[1]
        add(a, b)
        add(b, a)
    }
    // 以每一个节点为起点进行DFS遍历染色
    for i:=1; i<=n; i++ {
        // 如果已经染过色，则直接跳过
        if color[i] != 0 {
            continue
        }
        // 尝试将第一个节点从红色开始进行染色
        if (!dfs(i, 1)) {
            return false
        }
    }
    return true
}
```

### 专题总结
- [二分图思想](https://blog.csdn.net/Keven_11/article/details/122754558)
- [邻接表存储](https://mp.weixin.qq.com/s/2Ba8-NI7lQh2_MvMpg-CZg)
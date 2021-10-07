package leetcode_0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录每个元素的入度
	inDegree := make([]int, numCourses)
	// 用字典来标识一个图的数据结构, key 为先修课程，value 为后修课程
	graph := make(map[int][]int)

	// 首先求每个元素的初始入度
	for _, prerequisite := range prerequisites {
		inDegree[prerequisite[0]]++
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}

	// 将所有入度为 0 的课程入队
	queue := make([]int, 0)
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 记录所有已经可以修完的课程数
	count := 0
	// 进行 BFS
	for len(queue) > 0 {
		preCourse := queue[0]
		queue = queue[1:]
		count++

		// 将所有前序是该课程的入度都减一
		postCourse := graph[preCourse]
		for i := 0; i < len(postCourse); i++ {
			inDegree[postCourse[i]]--
			// 如果某个课程在减少入度后，入度变为0，则进队列
			if inDegree[postCourse[i]] == 0 {
				queue = append(queue, postCourse[i])
			}
		}
	}
	return count == numCourses
}

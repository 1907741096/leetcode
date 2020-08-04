package main

var classStatusList []int
var preClassList [][]int
var nextClassList [][]int
var preClassCountList []int
var sumCount int
var flag bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	flag = true
	sumCount = 0
	classStatusList = make([]int, numCourses) // 0-未遍历 1-遍历中 2-已遍历
	preClassList = make([][]int, numCourses) // 节点对应前置课程
	nextClassList = make([][]int, numCourses) // 节点对应后置课程
	preClassCountList = make([]int, numCourses) // 节点对应前置课程数目

	for _, p := range prerequisites {
		// 加入前置课程
		preClassList[p[0]] = append(preClassList[p[0]], p[1])
		// 加入后置课程
		nextClassList[p[1]] = append(nextClassList[p[1]], p[0])
		preClassCountList[p[0]] ++
	}
	//for _, p := range prerequisites {
	//	if !flag {
	//		break
	//	}
	//	// 将当前节点加入栈中
	//	dfs(p[0])
	//}

	for i, p := range preClassList {
		// 无前置节点，进入广度遍历
		if len(p) == 0 {
			bfs(i)
		}
	}
	if numCourses != sumCount {
		return false
	}

	return flag
}

// 深度优先
func dfs(nowClass int) {
	if !flag {
		return
	}
	if classStatusList[nowClass] == 0 {
		// 如果未遍历，当前节点设置为遍历中
		classStatusList[nowClass] = 1
	} else if classStatusList[nowClass] == 1 {
		// 如果遍历中，则出现环，没有拓扑排序，返回false
		flag = false
		return
	} else {
		// 如果已遍历，break
		return
	}
	// 遍历当前节点前置节点
	for _, nowPreClass := range preClassList[nowClass] {
		if classStatusList[nowPreClass] == 0 {
			// 未遍历
			dfs(nowPreClass)
		} else if classStatusList[nowPreClass] == 1 {
			// 出现环，没有拓扑排序，false
			flag = false
			return
		} else {
			// 已经遍历继续
		}
	}
	// 遍历完成，当前节点置为已完成
	classStatusList[nowClass] = 2
}

// 广度优先
func bfs(nowClass int) {
	sumCount ++
	queue := []int{nowClass}
	for len(queue) != 0 {
		nowClass = queue[0]
		queue = queue[1:]
		for _, v := range nextClassList[nowClass] {
			// 前置数-1
			preClassCountList[v] --
			if preClassCountList[v] == 0 {
				// 前置课全部学完时，加入队列
				queue = append(queue, v)
				// 完成课程节点+1
				sumCount ++
			}
		}
	}
}
package main

import (
	"fmt"

	"github.com/chentaihan/container/queue"
)

// 深度优先
func getProvince(citysConnected [][]int) int {
	citys := len(citysConnected)
	visited := make([]bool, citys)
	province := 0 // 计数器
	for i := 0; i < citys; i++ {
		if !visited[i] {
			// 深度优先
			dfs(i, citys, visited, citysConnected)
			province++
		}
	}
	return province
}

// 广度优先
func bfs(citysConnected [][]int) int {
	citys := len(citysConnected)
	visited := make([]bool, citys)
	province := 0 // 计数器
	queueLink := queue.NewQueueLink()
	for i := 0; i < citys; i++ {
		if !visited[i] {
			queueLink.Enqueue(i)
			for !queueLink.Empty() {
				k, _ := queueLink.Dequeue()
				visited[k.(int)] = true
				for j := 0; j < citys; j++ {
					if citysConnected[i][j] == 1 && !visited[j] {
						queueLink.Enqueue(j)
					}
				}
			}
			province++
		}
	}
	return province
}

// 并查集
func mergeFind(citysConnected [][]int) int {
	citys := len(citysConnected)
	head := make([]int, citys)
	level := make([]int, citys)
	for i := 0; i < citys; i++ {
		head[i] = i
		level[i] = i
	}
	for i := 0; i < citys; i++ {
		for j := i + 1; j < citys; j++ {
			if citysConnected[i][j] == 1 {
				merge(i, j, head, level)
			}
		}
	}
	count := 0
	for i := 0; i < citys; i++ {
		if head[i] == i {
			count++
		}
	}
	return count
}

func merge(x int, y int, head []int, level []int) {
	i := find(x, head)
	j := find(y, head)
	if i == j {
		return
	}
	if level[i] <= level[j] {
		head[i] = j
	} else {
		head[j] = i
	}
	if level[i] == level[j] {
		level[i]++
		level[j]++
	}
}

func find(x int, head []int) int {
	if head[x] == x {
		return x
	}
	head[x] = find(head[x], head)
	return head[x]
}

// 深度优先
func dfs(i int, citys int, visited []bool, citysConnected [][]int) {
	for j := 0; j < citys; j++ {
		if citysConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(j, citys, visited, citysConnected)
		}
	}
}

// 省份数量
func main() {
	fmt.Println(getProvince([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})) // 2
	fmt.Println(getProvince([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) // 3
	fmt.Println(bfs([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))         // 2
	fmt.Println(bfs([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))         // 3
	fmt.Println(mergeFind([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))   // 2
	fmt.Println(mergeFind([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))   // 3
}
